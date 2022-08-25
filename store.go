package nbc

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pmoieni/nimbus-cloud/internal/db/auth"
	"github.com/pmoieni/nimbus-cloud/internal/db/relation"
	"github.com/pmoieni/nimbus-cloud/internal/db/store"
	"github.com/pmoieni/nimbus-cloud/internal/req"
	"golang.org/x/exp/slices"
)

type StoreService struct {
	DBCon *sql.DB
}

type (
	permitUsersReq struct {
		Users []string `json:"users"`
	}
	getUserFilesRes struct {
		UserFiles      []store.File `json:"user_files"`
		PermittedFiles []store.File `json:"permitted_files"`
	}
)

const (
	mb = 1 << (10 * 2)
)

func (s *StoreService) Upload(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value(emailCtxKey).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// function body of a http.HandlerFunc
	r.Body = http.MaxBytesReader(w, r.Body, 100*mb+1024)
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// parse text field
	text := make([]byte, 512)
	p, err := reader.NextPart()
	// one more field to parse, EOF is considered as failure here
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if p.FormName() != "text" {
		http.Error(w, "text field is expected", http.StatusBadRequest)
		return
	}
	_, err = p.Read(text)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// parse file field
	p, err = reader.NextPart()
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if p.FormName() != "file" {
		http.Error(w, "file field is expected", http.StatusBadRequest)
		return
	}
	buf := bufio.NewReader(p)
	randID := uuid.New().String()
	randName := "./bucket/" + randID
	f, err := os.OpenFile(randName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	var maxSize int64 = 100 * mb
	lmt := io.MultiReader(buf, io.LimitReader(p, maxSize-511))
	written, err := io.Copy(f, lmt)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if written > maxSize {
		os.Remove(f.Name())
		http.Error(w, "file size over limit", http.StatusBadRequest)
		return
	}

	rawFileName := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, string(text))
	var fileName string
	var extension string
	if len(strings.TrimSpace(string(rawFileName))) > 0 {
		parts := strings.Split(p.FileName(), ".")
		if len(parts) > 0 {
			extension = strings.TrimSpace(parts[len(parts)-1])
			fileName = rawFileName + "." + extension

		} else {
			fileName = p.FileName()
		}
	} else {
		fileName = p.FileName()
	}

	q := auth.New(s.DBCon)
	user, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.saveFile(randID, fileName, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *StoreService) GetUserFiles(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value(emailCtxKey).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	q := auth.New(s.DBCon)
	user, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userFiles, err := s.getUserFiles(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	permittedFiles, err := s.getUserPermittedFiles(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := getUserFilesRes{
		UserFiles:      userFiles,
		PermittedFiles: permittedFiles,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *StoreService) PermitUsers(w http.ResponseWriter, r *http.Request) {
	on := chi.URLParam(r, "name")
	email, ok := r.Context().Value(emailCtxKey).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p := permitUsersReq{}
	if err := req.Parse(r, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uq := auth.New(s.DBCon)
	ownerInfo, err := uq.GetUserByEmail(context.Background(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fq := store.New(s.DBCon)

	fileInfo, err := fq.GetFileByObjectName(context.Background(), on)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	permittedUsers, err := uq.ListPermittedUsers(context.Background(), fileInfo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	puEmails := permittedUsers.GetEmails()
	if fileInfo.Owner != ownerInfo.ID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, email := range p.Users {
		if !slices.Contains(puEmails, email) && email != ownerInfo.Email {
			userInfo, err := uq.GetUserByEmail(context.Background(), email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			err = s.createUserFileRelation(fileInfo.ID, userInfo.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
}

// // difference returns the elements in `a` that aren't in `b`.
// func difference(a, b []string) []string {
// 	mb := make(map[string]struct{}, len(b))
// 	for _, x := range b {
// 		mb[x] = struct{}{}
// 	}
// 	var diff []string
// 	for _, x := range a {
// 		if _, found := mb[x]; !found {
// 			diff = append(diff, x)
// 		}
// 	}
// 	return diff
// }

func (s *StoreService) Download(w http.ResponseWriter, r *http.Request) {
	on := chi.URLParam(r, "name")
	fileInfo, ok := r.Context().Value(fileInfoCtxKey).(store.File)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename="+"\""+fileInfo.Name+"\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, "./bucket/"+on)
}

func (s *StoreService) Delete(w http.ResponseWriter, r *http.Request) {
	on := chi.URLParam(r, "name")
	err := s.removeFile(on)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.Remove("./bucket/" + on)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *StoreService) GetPermittedUsersEmails(w http.ResponseWriter, r *http.Request) {
	on := chi.URLParam(r, "name")
	fq := store.New(s.DBCon)
	fileInfo, err := fq.GetFileByObjectName(context.Background(), on)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uq := auth.New(s.DBCon)
	users, err := uq.ListPermittedUsers(context.Background(), fileInfo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := usersEmailListRes{}
	for _, user := range users {
		res.Users = append(res.Users, user.Email)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ------------------------- SERVICE -----------------------------
func (s *StoreService) saveFile(rn, fn string, uID int64) error {
	q := store.New(s.DBCon)
	file := store.CreateFileParams{}
	file.Name = fn
	file.ObjectName = rn
	file.Owner = uID
	_, err := q.CreateFile(context.Background(), &file)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoreService) removeFile(on string) error {
	q := store.New(s.DBCon)
	err := q.DeleteFileByObjectName(context.Background(), on)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoreService) getUserFiles(uID int64) ([]store.File, error) {
	ctx := context.Background()
	q := store.New(s.DBCon)
	files, err := q.ListOwnerFiles(ctx, uID)
	if err != nil {
		return []store.File{}, err
	}

	return files, nil
}

func (s *StoreService) getUserPermittedFiles(uID int64) ([]store.File, error) {
	ctx := context.Background()
	q := store.New(s.DBCon)
	files, err := q.ListPermittedFiles(ctx, uID)
	if err != nil {
		return []store.File{}, err
	}

	return files, nil
}

func (s *StoreService) createUserFileRelation(fID, uID int64) error {
	q := relation.New(s.DBCon)
	r := relation.CreateUserFileRelationParams{}
	r.FileID = fID
	r.UserID = uID

	_, err := q.CreateUserFileRelation(context.Background(), &r)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoreService) removeUserRelation(uID int64) error {
	q := relation.New(s.DBCon)

	err := q.DeleteRelationByUserID(context.Background(), uID)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoreService) removeFileRelation(fID int64) error {
	q := relation.New(s.DBCon)

	err := q.DeleteRelationByFileID(context.Background(), fID)
	if err != nil {
		return err
	}

	return nil
}

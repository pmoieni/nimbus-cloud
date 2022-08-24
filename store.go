package nbc

import (
	"bufio"
	"context"
	"database/sql"
	"io"
	"net/http"
	"os"
	"time"

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
		Files []string `json:"files"`
		Users []string `json:"users"`
	}
	getUserFilesReq struct {
		UserID int64 `json:"user_id"`
	}
	getUserFilesRes struct {
		UserFiles      []store.File `json:"user_files"`
		PermittedFiles []store.File `json:"permitted_files"`
	}
	downloadReq struct {
		ObjectName string `json:"object_name"`
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
	randName := uuid.New().String()
	f, err := os.OpenFile("./bucket/"+randName, os.O_WRONLY|os.O_CREATE, 0666)
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

	q := auth.New(s.DBCon)
	user, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// f.Seek(0, io.SeekStart) // seek reader to the start to prevent EOF
	err = s.saveFile(randName, p.FileName(), user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *StoreService) GetUserFiles(w http.ResponseWriter, r *http.Request) {
	p := getUserFilesReq{}
	if err := req.Parse(r, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userFiles, err := s.getUserFiles(p.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	permittedFiles, err := s.getUserPermittedFiles(p.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := getUserFilesRes{
		UserFiles:      userFiles,
		PermittedFiles: permittedFiles,
	}

	w.WriteHeader(http.StatusOK)
	req.WriteJSON(w, res)
}

func (s *StoreService) PermitUsers(w http.ResponseWriter, r *http.Request) {
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

	for _, fileON := range p.Files {
		fileInfo, err := fq.GetFileByObjectName(context.Background(), fileON)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if fileInfo.Owner != ownerInfo.ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fileUsers, err := uq.ListPermittedUsers(context.Background(), fileInfo.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		usersEmails := fileUsers.GetEmails()
		addedUsers := difference(p.Users, usersEmails)
		for _, email := range p.Users {
			userInfo, err := uq.GetUserByEmail(context.Background(), email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if slices.Contains(addedUsers, userInfo.Email) {
				err = s.createUserFileRelation(fileInfo.ID, userInfo.ID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
				err = s.removeUserRelation(userInfo.ID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
	}

	w.WriteHeader(http.StatusOK)
}

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func (s *StoreService) Download(w http.ResponseWriter, r *http.Request) {
	on := chi.URLParam(r, "name")
	fileName, ok := r.Context().Value(fileNameCtxKey).(string)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(w, r, "./bucket/"+on)
}

func (s *StoreService) Delete(w http.ResponseWriter, r *http.Request) {

}

func (s *StoreService) Purge(w http.ResponseWriter, r *http.Request) {

}

// ------------------------- SERVICE -----------------------------
func (s *StoreService) saveFile(rn, fn string, uID int64) error {
	ctx := context.Background()
	q := store.New(s.DBCon)
	file := store.CreateFileParams{
		Name:       fn,
		ObjectName: rn,
		Owner:      uID,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}
	_, err := q.CreateFile(ctx, &file)
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
	r.CreatedAt = time.Now().UTC()
	r.UpdatedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}

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

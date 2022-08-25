package nbc

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/pmoieni/nimbus-cloud/internal/db/auth"
	"github.com/pmoieni/nimbus-cloud/internal/db/store"
)

var emailCtxKey = ctxKey("email")

func (s *AuthService) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at := r.Header.Get(authorizationHeader)
		bearer := strings.Split(at, " ")
		if !(len(bearer) > 1) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userInfo, err := parseAccessTokenWithValidate(bearer[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		privateClaims := userInfo.PrivateClaims()
		email, ok := privateClaims["email"].(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), emailCtxKey, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

var fileInfoCtxKey = ctxKey("file-info")

func (s *StoreService) CheckFilePermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email, ok := r.Context().Value(emailCtxKey).(string)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		on := chi.URLParam(r, "name")
		fq := store.New(s.DBCon)
		fileInfo, err := fq.GetFileByObjectName(context.Background(), on)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		uq := auth.New(s.DBCon)
		userInfo, err := uq.GetUserByEmail(context.Background(), email)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), fileInfoCtxKey, fileInfo)
		if fileInfo.Owner != userInfo.ID {
			permittedFiles, err := fq.ListPermittedFiles(context.Background(), userInfo.ID)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			for _, file := range permittedFiles {
				if file.ObjectName == on {
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *StoreService) CheckFileOwner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email, ok := r.Context().Value(emailCtxKey).(string)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		on := chi.URLParam(r, "name")
		fq := store.New(s.DBCon)
		fileInfo, err := fq.GetFileByObjectName(context.Background(), on)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		uq := auth.New(s.DBCon)
		userInfo, err := uq.GetUserByEmail(context.Background(), email)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), fileInfoCtxKey, fileInfo)
		if fileInfo.Owner != userInfo.ID {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

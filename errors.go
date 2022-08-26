package nbc

import (
	"log"
	"net/http"

	"github.com/pmoieni/nimbus-cloud/internal/req"
)

var (
	errInvalidUpdateInfo   = req.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid Username"}
	errInvalidRegisterInfo = req.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid Register Info"}
	errUserAlreadyExists   = req.ErrorResponse{Status: http.StatusForbidden, Message: "User with the same email already exists."}
	errUserNotFound        = req.ErrorResponse{Status: http.StatusNotFound, Message: "User not found."}
	errBadPassword         = req.ErrorResponse{Status: http.StatusBadRequest, Message: "Password must contain at least 8 characters and one number"}
	errInvalidEmail        = req.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid Email"}
	errNoCookie            = req.ErrorResponse{Status: http.StatusUnauthorized, Message: "Cookie not found."}
)

func handlerError(w http.ResponseWriter, err error) {
	if err != nil {
		if httpError, ok := err.(*req.ErrorResponse); ok {
			http.Error(w, httpError.Message, httpError.Status)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

package nbc

import (
	"net/http"

	"github.com/pmoieni/nimbus-cloud/internal/req"
)

var (
	errInvalidUpdateInfo   = req.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid Username"}
	errInvalidRegisterInfo = req.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid Register Info"}
	errUserAlreadyExists   = req.ErrorResponse{Status: http.StatusForbidden, Message: "User with the same email already exists."}
	errUserNotFound        = req.ErrorResponse{Status: http.StatusNotFound, Message: "User not found."}
	errBadPassword         = req.ErrorResponse{Status: http.StatusBadRequest, Message: "Bad password"}
	errInvalidEmail        = req.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid Email"}
	errNoCookie            = req.ErrorResponse{Status: http.StatusUnauthorized, Message: "Cookie not found."}
)

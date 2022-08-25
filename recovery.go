package nbc

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/go-redis/redis/v9"
// 	"github.com/pmoieni/nimbus-cloud/internal/req"
// 	"golang.org/x/crypto/bcrypt"
// )

// type resetPasswordReq struct {
// 	Email       string `json:"email"`
// 	Token       string `json:"token"`
// 	NewPassword string `json:"new_password"`
// }

// func (s *AuthService) ResetPassword(w http.ResponseWriter, r *http.Request) {
// 	rp := resetPasswordReq{}
// 	if err := req.Parse(r, &rp); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err := s.resetPassword(&rp)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	return c.JSON(http.StatusOK, "success")
// }

// func (s *AuthService) SendPasswordResetToken(w http.ResponseWriter, r *http.Request) {

// }

// // ----------------------- SERVICES ------------------------

// func (s *AuthService) resetPassword(r *resetPasswordReq) error {
// 	getTokenRedisDTO := store.RedisDTO{
// 		Key: r.Token,
// 	}
// 	v, err := getTokenRedisDTO.Get(store.PasswordResetTokenDB)
// 	if err != nil {
// 		switch err {
// 		case redis.Nil:
// 			return &req.ErrorResponse{Status: http.StatusForbidden, Message: http.StatusText(http.StatusForbidden)}
// 		default:
// 			return err
// 		}
// 	}

// 	// remove token since it's already used
// 	s.RedisPasswordTokenDB.Del(r.Token)

// 	tInfo := PasswordResetTokenReq{}
// 	err = json.Unmarshal([]byte(v), &tInfo)
// 	if err != nil {
// 		return err
// 	}

// 	// check if the token belongs to the email which is trying to reset the token
// 	if r.Email != tInfo.Email {
// 		return &models.ErrorResponse{Status: http.StatusForbidden, Message: http.StatusText(http.StatusForbidden)}
// 	}

// 	u := store.User{
// 		Email: r.Email,
// 	}
// 	userInfo, err := u.GetUser()
// 	if err != nil {
// 		return err
// 	}

// 	// check if new password is valid
// 	err = validatePassword(r.NewPassword)
// 	if err != nil {
// 		return &models.ErrorResponse{Status: http.StatusForbidden, Message: errBadPassword}
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.NewPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	updateUserInfo := store.User{
// 		Password: string(hashedPassword),
// 	}

// 	err = userInfo.Update(&updateUserInfo)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

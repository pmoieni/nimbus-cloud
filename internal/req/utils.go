package req

import (
	"encoding/json"
	"net/http"
)

func Parse(r *http.Request, out interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&out)
	if err != nil {
		return err
	}

	return nil
}

type ErrorResponse struct {
	Status  int
	Message string
}

// custom error type for detecting internal application errors
func (e *ErrorResponse) Error() string {
	return e.Message
}

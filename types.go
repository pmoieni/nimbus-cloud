package nbc

type ctxKey string // use a type of type string to make sure ctx key is coming from middleware

type usersEmailListRes struct {
	Users []string `json:"users"`
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package auth

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64        `db:"id" json:"id"`
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`
}
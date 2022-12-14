// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: auth_query.sql

package auth

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
        username,
        email,
        password,
        created_at,
        updated_at,
        deleted_at
    )
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	Username  string       `db:"username" json:"username"`
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) (sql.Result, error) {
	return q.exec(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
	)
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password, created_at, updated_at, deleted_at
FROM users
WHERE email = ?
LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, password, created_at, updated_at, deleted_at
FROM users
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, password, created_at, updated_at, deleted_at
FROM users
ORDER BY id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :execresult
UPDATE users
    SET username = ?
    WHERE id = ?
`

type UpdateUserParams struct {
	Username string `db:"username" json:"username"`
	ID       int64  `db:"id" json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg *UpdateUserParams) (sql.Result, error) {
	return q.exec(ctx, q.updateUserStmt, updateUser, arg.Username, arg.ID)
}

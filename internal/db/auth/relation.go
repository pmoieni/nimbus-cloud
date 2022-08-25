package auth

import (
	"context"
)

var listPermittedUsers = `SELECT 
	u.id,
	username,
    email
	FROM user_file_relation uf
    INNER JOIN users u ON u.id = uf.user_id
	WHERE uf.file_id = ?;`

type Users []User

func (q *Queries) ListPermittedUsers(ctx context.Context, fID int64) (Users, error) {
	stmt, err := q.db.PrepareContext(context.Background(), listPermittedUsers)
	if err != nil {
		return nil, err
	}
	rows, err := q.query(ctx, stmt, listPermittedUsers, fID)
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

func (users Users) GetEmails() []string {
	var data []string
	for _, user := range users {
		data = append(data, user.Email)
	}
	return data
}

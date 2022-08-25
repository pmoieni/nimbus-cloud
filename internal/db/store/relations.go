package store

import "context"

var listSharedFiles = `SELECT
	f.id,
	name,
	object_name,
	owner
	FROM user_file_relation uf INNER JOIN files f ON f.id = uf.file_id WHERE uf.user_id = ?`

func (q *Queries) ListPermittedFiles(ctx context.Context, uID int64) ([]File, error) {
	stmt, err := q.db.PrepareContext(context.Background(), listSharedFiles)
	if err != nil {
		return nil, err
	}
	rows, err := q.query(ctx, stmt, listSharedFiles, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []File{}
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ObjectName,
			&i.Owner,
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

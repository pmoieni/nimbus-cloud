-- name: GetFileByID :one
SELECT *
FROM files
WHERE id = ?
LIMIT 1;

-- name: GetFileByName :one
SELECT *
FROM files
WHERE name = ?
LIMIT 1;

-- name: GetFileByObjectName :one
SELECT *
FROM files
WHERE object_name = ?
LIMIT 1;

-- name: ListFiles :many
SELECT *
FROM files
ORDER BY id;

-- name: ListOwnerFiles :many
SELECT *
FROM files
WHERE owner = ?
ORDER BY id;

-- name: CreateFile :execresult
INSERT INTO files (
        name,
        object_name,
        owner,
        created_at,
        updated_at,
        deleted_at
    )
VALUES (?, ?, ?, ?, ?, ?);

-- name: DeleteFileByID :exec
DELETE FROM files
WHERE id = ?;

-- name: DeleteFileByObjectName :exec
DELETE FROM files
WHERE object_name = ?;
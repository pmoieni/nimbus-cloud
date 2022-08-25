-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ?
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO users (
        email,
        password,
        created_at,
        updated_at,
        deleted_at
    )
VALUES (?, ?, ?, ?, ?);

-- name: UpdateUser :execresult
UPDATE users
    SET email = ?,
    password = ?
    WHERE id = ?;

-- name: UpdateUserEmail :execresult
UPDATE users
    SET email = ?
    WHERE id = ?;

-- name: UpdateUserPassword :execresult
UPDATE users
    SET email = ?
    WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

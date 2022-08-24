-- name: GetFileUserRelationByUserID :one
SELECT *
FROM user_file_relation
WHERE user_id = ?
LIMIT 1;

-- name: GetFileUserRelationByFileID :one
SELECT *
FROM user_file_relation
WHERE file_id = ?
LIMIT 1;

-- name: CreateUserFileRelation :execresult
INSERT INTO user_file_relation (
        file_id,
        user_id,
        created_at,
        updated_at,
        deleted_at
    )
VALUES (?, ?, ?, ?, ?);

-- name: DeleteRelationByUserID :exec
DELETE FROM user_file_relation
WHERE user_id = ?;

-- name: DeleteRelationByFileID :exec
DELETE FROM user_file_relation
WHERE file_id = ?;

-- -- name: ListFileUsers :many
-- SELECT u.id,
--     email
-- FROM user_file_relation uf
--     INNER JOIN users u ON u.id = uf.user_id
-- WHERE uf.file_id = ?;

-- -- name: ListUserFiles :many
-- SELECT f.id,
--     name,
--     object_name,
--     owner
-- FROM user_file_relation uf
--     INNER JOIN files f ON f.id = uf.file_id
-- WHERE uf.user_id = ?;

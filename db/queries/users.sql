-- name: CreateUser :one
INSERT INTO "users" (
    username,
    email,
    password,
    avatar,
    role
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE username = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users 
SET 
    username = COALESCE(sqlc.narg(username), username),
    email = COALESCE(sqlc.narg(email), email),
    password = COALESCE(sqlc.narg(password), password),
    avatar = COALESCE(sqlc.narg(avatar), avatar),
    role = COALESCE(sqlc.narg(role), role)
WHERE 
    username = sqlc.arg(username)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM users
WHERE username = $1;
-- name: CreateSubleddit :one
INSERT INTO "subleddits" (
    name,
    user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: DeleteSubleddit :exec
DELETE FROM "subleddits"
WHERE name = $1;
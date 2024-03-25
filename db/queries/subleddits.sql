-- name: CreateSubleddit :one
WITH Subleddit AS (
    INSERT INTO "subleddits" (
        name,
        user_id
    ) VALUES (
        $1, $2
    ) RETURNING *
) SELECT Subleddit.*, sqlc.embed(users) FROM Subleddit JOIN users ON Subleddit.user_id = users.id;

-- name: GetSubleddit :one
SELECT subleddits.*, sqlc.embed(users)
FROM subleddits
JOIN users ON users.id = subleddits.user_id
WHERE subleddits.name = $1;

-- name: DeleteSubleddit :exec
DELETE FROM "subleddits"
WHERE name = $1;
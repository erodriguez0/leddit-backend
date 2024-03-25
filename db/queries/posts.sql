-- name: CreatePost :one
WITH Post AS (
    INSERT INTO "posts" (
        title,
        url,
        body,
        subleddit_id,
        user_id
    ) VALUES (
        $1, $2, $3, $4, $5
    ) RETURNING *
) 
SELECT Post.*, sqlc.embed(users), sqlc.embed(subleddits) 
FROM Post
JOIN users ON Post.user_id = users.id
JOIN subleddits ON Post.subleddit_id = subleddits.id;
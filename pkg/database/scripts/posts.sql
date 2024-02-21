-- name: GetPosts :many
SELECT * FROM posts;

-- name: GetPostById :one
SELECT * FROM posts WHERE id = $1;

-- name: CreatePost :exec
INSERT INTO posts(title, content, created_at, updated_at)
VALUES ($1, $2, $3, $4);

-- name: DeletePostById :exec
DELETE FROM posts WHERE id = $1;

-- name: UpdatePostById :one
UPDATE posts
SET title=$1, content=$2, updated_at=$3
WHERE id = $4
RETURNING *;
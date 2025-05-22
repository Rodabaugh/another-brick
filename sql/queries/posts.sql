-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, content)
VALUES (
    gen_random_uuid(), NOW(), NOW(), $1
)
RETURNING *;

-- name: GetPosts :many
SELECT * FROM posts;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1;

-- name: DeletePostByID :exec
DELETE FROM posts WHERE id = $1;
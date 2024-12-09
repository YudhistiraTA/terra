-- name: FuzzySearchPosts :many
SELECT *
FROM posts
WHERE 
  posts.user_id = @user_id 
  AND (sqlc.narg('search_term')::text IS NULL OR title % sqlc.narg('search_term')::text OR content % sqlc.narg('search_term')::text)
  AND (sqlc.narg('cursor')::text IS NULL OR created_at <= (SELECT created_at FROM posts WHERE id::text = sqlc.narg('cursor')::text))
ORDER BY
    created_at DESC
LIMIT 6;

-- name: GetPreviousCursor :one
SELECT id
FROM posts
WHERE 
  posts.user_id = @user_id
  AND (sqlc.narg('search_term')::text IS NULL OR title % sqlc.narg('search_term')::text OR content % sqlc.narg('search_term')::text)
  AND created_at > (SELECT created_at FROM posts WHERE posts.id::text = @cursor)
ORDER BY created_at ASC
LIMIT 5
offset 4;

-- name: GetPostById :one
SELECT *
FROM posts
WHERE id = $1;

-- name: ListPosts :many
SELECT *
FROM posts
WHERE user_id = $1
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreatePost :exec
INSERT INTO posts (title, content, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdatePost :exec
UPDATE posts
SET title = $1, content = $2
WHERE id = $3
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
RETURNING *;
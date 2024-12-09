-- name: FuzzySearchPosts :many
SELECT *
FROM posts
WHERE 
  posts.user_id = @user_id 
  AND (sqlc.narg('search_term')::text IS NULL OR title % sqlc.narg('search_term')::text OR content % sqlc.narg('search_term')::text)
  AND (sqlc.narg('cursor')::text IS NULL OR created_at <= (SELECT created_at FROM posts WHERE id::text = sqlc.narg('cursor')::text))
ORDER BY
  CASE
    WHEN sqlc.narg('search_term')::text IS NOT NULL AND title % sqlc.narg('search_term')::text THEN similarity(title, sqlc.narg('search_term')::text)
    ELSE 0
  END DESC,
  CASE
    WHEN sqlc.narg('search_term')::text IS NOT NULL AND content % sqlc.narg('search_term')::text THEN similarity(content, sqlc.narg('search_term')::text)
    ELSE 0
  END DESC,
  CASE
    WHEN sqlc.narg('search_term')::text IS NULL THEN created_at
  END DESC
LIMIT 11;

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
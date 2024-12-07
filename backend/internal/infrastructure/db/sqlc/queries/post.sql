-- name: FuzzySearchPosts :many
SELECT *
FROM posts
WHERE title % 'search_term' OR content % 'search_term'
ORDER BY
  CASE
    WHEN title % 'search_term' THEN similarity(title, 'search_term')
    ELSE 0
  END DESC,
  CASE
    WHEN content % 'search_term' THEN similarity(content, 'search_term')
    ELSE 0
  END DESC
LIMIT 10;

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
-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateRefreshById :exec
UPDATE users
SET refresh_token = $2
WHERE email = $1;
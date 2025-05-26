-- name: UserGetById :one
SELECT * 
FROM users
WHERE id = $1;

-- name: UserCreate :one
INSERT INTO users (username) VALUES ($1) RETURNING id;

-- name: UserUpdate :execrows
UPDATE users
SET
username = COALESCE(sqlc.narg('username'), username),
status = COALESCE(sqlc.narg('status'), status)
WHERE id = $1;

-- name: UserDelete :execrows
DELETE FROM users
WHERE id = $1;

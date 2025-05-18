-- name: UserGetById :one
SELECT * 
FROM users
WHERE id = $1;

-- name: UserCreate :one
INSERT INTO users DEFAULT VALUES RETURNING id;

-- name: UserUpdate :execrows
UPDATE users
SET
username = COALESCE(sqlc.narg('username'), username)
WHERE id = $1;

-- name: UserDelete :execrows
DELETE FROM users
WHERE id = $1;

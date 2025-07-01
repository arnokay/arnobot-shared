-- name: CoreUserCommandGetByUserID :many
SELECT
    *
FROM
    core.user_commands
WHERE
    user_id = $1
ORDER BY
    updated_at DESC;

-- name: CoreUserCommandGetOne :one
SELECT
    *
FROM
    core.user_commands
WHERE
    user_id = $1
    AND name = $2;

-- name: CoreUserCommandCreate :one
INSERT INTO core.user_commands (user_id, name, text, reply)
    VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: CoreUserCommandUpdate :one
UPDATE
    core.user_commands
SET
    name = COALESCE(sqlc.narg ('NewName')::varchar(50), name),
    text = COALESCE(sqlc.narg ('text')::text, text),
    reply = COALESCE(sqlc.narg ('reply')::bool, reply)
WHERE
    user_id = $1
    AND name = $2
RETURNING
    *;

-- name: CoreUserCommandDelete :one
DELETE FROM core.user_commands
WHERE user_id = $1
    AND name = $2
RETURNING
    *;


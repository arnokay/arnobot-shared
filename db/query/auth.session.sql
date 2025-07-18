-- name: AuthSessionValidate :one
UPDATE
    auth.sessions
SET
    last_used_at = CURRENT_TIMESTAMP
WHERE
    token = $1
RETURNING
    status;

-- name: AuthSessionDelete :one
DELETE FROM auth.sessions
WHERE token = $1
RETURNING
    *;

-- name: AuthSessionOldDeactivate :exec
UPDATE
    auth.sessions
SET
    status = 'disabled'
WHERE
    auth.sessions.user_id = $1
    AND last_used_at < (
        SELECT
            last_used_at
        FROM
            auth.sessions
        WHERE
            user_id = $1
        ORDER BY
            last_used_at DESC
        LIMIT 1 OFFSET $2);

-- name: AuthSessionGet :one
SELECT
    *
FROM
    auth.sessions
WHERE
    token = $1;

-- name: AuthSessionCreate :one
INSERT INTO auth.sessions (user_id)
    VALUES ($1)
RETURNING
    *;

-- name: AuthSessionActiveCount :one
SELECT
    count(token)
FROM
    auth.sessions
WHERE
    user_id = $1
    AND status = 'active';

-- name: AuthSessionActiveGet :many
SELECT
    *
FROM
    auth.sessions
WHERE
    user_id = $1
    AND status = 'active';

-- name: AuthSessionGetOwner :one
SELECT
    sqlc.embed(pUsers)
FROM
    auth.sessions
    LEFT JOIN public.users AS pUsers ON auth.sessions.user_id = pUsers.id
WHERE
    token = $1;


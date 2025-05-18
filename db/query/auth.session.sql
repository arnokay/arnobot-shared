-- name: AuthSessionValidate :one
UPDATE auth.sessions
SET
last_seen_at = CURRENT_TIMESTAMP
WHERE token = $1
RETURNING status;

-- name: AuthSessionGet :one
SELECT *
FROM auth.sessions
WHERE token = $1;

-- name: AuthSessionCreate :one
INSERT INTO auth.sessions (
  user_id
) VALUES (
  $1
) RETURNING *;

-- name: AuthSessionActiveCount :one
SELECT count(token)
FROM auth.sessions
WHERE user_id = $1;

-- name: AuthSessionActiveGet :many
SELECT *
FROM auth.sessions
WHERE user_id = $1 AND status = 'active';

-- name: AuthSessionGetOwner :one
SELECT sqlc.embed(pUsers)
FROM auth.sessions
LEFT JOIN public.users as pUsers ON auth.sessions.user_id = public.users.id
WHERE token = $1;

-- name: CoreUserScheduledMessageGetByUserID :many
SELECT
    *
FROM
    core.user_scheduled_messages
WHERE
    user_id = $1
ORDER BY
    updated_at DESC;

-- name: CoreUserScheduledMessageGetOne :one
SELECT
    *
FROM
    core.user_scheduled_messages
WHERE
    id = $1;

-- name: CoreUserScheduledMessageCreate :one
INSERT INTO core.user_scheduled_messages (user_id, text, cooldown, platforms)
    VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: CoreUserScheduledMessageUpdate :one
UPDATE
    core.user_scheduled_messages
SET
    text = COALESCE(sqlc.narg('text')::text, text),
    cooldown = COALESCE(sqlc.narg('cooldown')::bigint, cooldown),
    platforms = COALESCE(sqlc.narg('platforms')::public.platform[], platforms)
WHERE
    id = $1
RETURNING
    *;

-- name: CoreUserScheduledMessageDelete :one
DELETE FROM core.user_scheduled_messages
WHERE id = $1
RETURNING
    *;


-- name: KickBotCreate :one
INSERT INTO kick.bots (user_id, broadcaster_id, bot_id)
    VALUES ($1, $2, $3)
RETURNING
    *;

-- name: KickBotGet :one
SELECT
    *
FROM
    kick.bots
WHERE
    user_id = $1
    AND bot_id = $2;

-- name: KickBotsGet :many
SELECT
    *
FROM
    kick.bots
WHERE (sqlc.narg ('user_id')::uuid IS NULL
    OR user_id = sqlc.narg ('user_id'))
AND (sqlc.narg ('broadcaster_id')::int IS NULL
    OR broadcaster_id = sqlc.narg ('broadcaster_id'))
AND (sqlc.narg ('bot_id')::int IS NULL
    OR bot_id = sqlc.narg ('bot_id'));

-- name: KickBotDelete :execrows
DELETE FROM kick.bots
WHERE user_id = $1;


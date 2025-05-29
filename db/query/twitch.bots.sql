-- name: TwitchBotCreate :one
INSERT INTO twitch.bots (
  user_id,
  broadcaster_id,
  bot_id
) VALUES (
  $1,
  $2,
  $3
) RETURNING *;

-- name: TwitchBotGet :one
SELECT *
FROM twitch.bots
WHERE
user_id = $1 AND bot_id = $2;

-- name: TwitchBotsGet :many
SELECT *
FROM twitch.bots
WHERE 
(sqlc.narg('user_id')::int IS NULL OR user_id = sqlc.narg('user_id')) AND
(sqlc.narg('broadcaster_id')::text IS NULL OR broadcaster_id = sqlc.narg('broadcaster_id')) AND
(sqlc.narg('bot_id')::text IS NULL OR bot_id = sqlc.narg('bot_id'));

-- name: TwitchBotDelete :execrows
DELETE FROM twitch.bots
WHERE user_id = $1;

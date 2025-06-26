-- name: KickDefaultBotUpdate :execrows
INSERT INTO kick.default_bot (main, bot_id)
VALUES (true, $1)
ON CONFLICT (main) DO UPDATE
  SET bot_id = $1;

-- name: KickDefaultBotGet :one
SELECT *
FROM kick.default_bot
WHERE main = true;


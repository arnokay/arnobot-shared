-- name: TwitchDefaultBotUpdate :execrows
INSERT INTO twitch.default_bot (main, bot_id)
VALUES (true, $1)
ON CONFLICT (main) DO UPDATE
  SET bot_id = $1;

-- name: TwitchDefaultBotGet :one
SELECT *
FROM twitch.default_bot
WHERE main = true;

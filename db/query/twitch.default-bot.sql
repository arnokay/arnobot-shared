-- name: TwitchDefaultBotUpdate :execrows
INSERT INTO twitch.default_bot (main, twitch_user_id)
VALUES (true, $1)
ON CONFLICT (main) DO UPDATE
  SET twitch_user_id = $1;

-- name: TwitchDefaultBotGet :one
SELECT *
FROM twitch.default_bot
WHERE main = true;

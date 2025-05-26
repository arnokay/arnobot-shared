-- name: TwitchBotCreate :one
INSERT INTO twitch.bots (
  user_id,
  twitch_user_id 
) VALUES (
  $1,
  $2
) RETURNING user_id;

-- name: TwitchBotDelete :execrows
DELETE FROM twitch.bots
WHERE user_id = $1;

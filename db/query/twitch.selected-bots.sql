-- name: TwitchSelectedBotChange :one
INSERT INTO twitch.selected_bots (user_id, bot_id, broadcaster_id)
VALUES ($1, $2, $3)
ON CONFLICT (user_id) DO UPDATE
  SET 
  bot_id = $2,
  broadcaster_id = $3,
  updated_at = CURRENT_TIMESTAMP
  RETURNING *;

-- name: TwitchSelectedBotGetByUserID :one
SELECT *
FROM twitch.selected_bots
WHERE user_id = $1;

-- name: TwitchSelectedBotGetByBroadcasterID :one
SELECT *
FROM twitch.selected_bots
WHERE broadcaster_id = $1;

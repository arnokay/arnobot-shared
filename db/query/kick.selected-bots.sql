-- name: KickSelectedBotChange :one
INSERT INTO kick.selected_bots (user_id, bot_id, broadcaster_id)
VALUES ($1, $2, $3)
ON CONFLICT (user_id) DO UPDATE
  SET 
  bot_id = $2,
  broadcaster_id = $3,
  updated_at = CURRENT_TIMESTAMP
  RETURNING *;

-- name: KickSelectedBotGetByUserID :one
SELECT *
FROM kick.selected_bots
WHERE user_id = $1;

-- name: KickSelectedBotGetByBroadcasterID :one
SELECT *
FROM kick.selected_bots
WHERE broadcaster_id = $1;

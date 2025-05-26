-- name: TwitchSelectedBotCreateOrUpdate :one
INSERT INTO twitch.selected_bots (user_id, twitch_user_id)
VALUES ($1, $2)
ON CONFLICT (user_id) DO UPDATE
  SET 
  twitch_user_id = $2,
  updated_at = CURRENT_TIMESTAMP
RETURNING user_id;

-- name: TwitchSelectedBotGet :one
SELECT *
FROM twitch.selected_bots
WHERE user_id = $1;

-- name: KickSelectedBotChange :one
INSERT INTO kick.selected_bots (user_id, bot_id, broadcaster_id, enabled)
    VALUES ($1, $2, $3, $4)
ON CONFLICT (user_id)
    DO UPDATE SET
        bot_id = $2,
        broadcaster_id = $3,
        enabled = $4,
        updated_at = CURRENT_TIMESTAMP
    RETURNING
        *;

-- name: KickSelectedBotStatusChange :one
UPDATE
    twitch.selected_bots
SET
    enabled = $2
WHERE
    user_id = $1
RETURNING
    *;

-- name: KickSelectedBotGetByUserID :one
SELECT
    *
FROM
    kick.selected_bots
WHERE
    user_id = $1;

-- name: KickSelectedBotGetByBroadcasterID :one
SELECT
    *
FROM
    kick.selected_bots
WHERE
    broadcaster_id = $1;


-- name: TwitchSelectedBotChange :one
INSERT INTO twitch.selected_bots (user_id, bot_id, broadcaster_id, enabled)
    VALUES ($1, $2, $3, $4)
ON CONFLICT (user_id)
    DO UPDATE SET
        bot_id = $2,
        broadcaster_id = $3,
        enabled = $4,
        updated_at = CURRENT_TIMESTAMP
    RETURNING
        *;

-- name: TwitchSelectedBotStatusChange :one
UPDATE
    twitch.selected_bots
SET
    enabled = $2
WHERE
    user_id = $1
RETURNING
    *;

-- name: TwitchSelectedBotGetByUserID :one
SELECT
    *
FROM
    twitch.selected_bots
WHERE
    user_id = $1;

-- name: TwitchSelectedBotGetByBroadcasterID :one
SELECT
    *
FROM
    twitch.selected_bots
WHERE
    broadcaster_id = $1;


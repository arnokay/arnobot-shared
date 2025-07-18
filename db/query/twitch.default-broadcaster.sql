-- name: TwitchDefaultBroadcasterUpdate :one
INSERT INTO twitch.default_broadcaster (main, broadcaster_id)
    VALUES (TRUE, $1)
ON CONFLICT (main)
    DO UPDATE SET
        broadcaster_id = $1
    RETURNING
        broadcaster_id;

-- name: TwitchDefaultBroadcasterGet :one
SELECT
    broadcaster_id
FROM
    twitch.default_broadcaster
WHERE
    main = TRUE;


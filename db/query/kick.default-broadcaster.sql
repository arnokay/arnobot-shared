-- name: KickDefaultBroadcasterUpdate :execrows
INSERT INTO kick.default_broadcaster (main, broadcaster_id)
    VALUES (TRUE, $1)
ON CONFLICT (main)
    DO UPDATE SET
        broadcaster_id = $1
    RETURNING
        broadcaster_id;

-- name: KickDefaultBroadcasterGet :one
SELECT
    broadcaster_id
FROM
    kick.default_broadcaster
WHERE
    main = TRUE;


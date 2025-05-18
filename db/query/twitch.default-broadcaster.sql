-- name: DefaultBroadcasterUpdate :execrows
INSERT INTO twitch.default_broadcaster (main, twitch_user_id)
VALUES (true, $1)
ON CONFLICT (main) DO UPDATE
  SET twitch_user_id = $1;

-- name: DefaultBroadcasterGet :one
SELECT *
FROM twitch.default_broadcaster
WHERE main = true;

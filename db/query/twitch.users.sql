-- name: TwitchUserCreate :one
INSERT INTO twitch.users (
id,
username,
display_name,
type,
broadcaster_type,
profile_image_url,
created_at
) VALUES (
$1,
$2,
$3,
$4,
$5,
$6,
$7
) RETURNING id;

-- name: TwitchUserConnectProvider :execrows
UPDATE twitch.users
SET
auth_provider_id = $1
WHERE id = $2;

-- name: TwitchWebhookCreate :one
INSERT INTO twitch.webhooks (
  subscription_id,
  subscription_status,
  event,
  callback,
  user_id,
  broadcaster_id,
  bot_id
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
) RETURNING subscription_id;

-- name: TwitchWebhookUpdateStatus :execrows
UPDATE twitch.webhooks
SET
status = COALESCE($2, status),
subscription_status = COALESCE($3, status),
updated_at = CURRENT_TIMESTAMP
WHERE subscription_id = $1;

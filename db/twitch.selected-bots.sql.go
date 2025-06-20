// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: twitch.selected-bots.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const twitchSelectedBotChange = `-- name: TwitchSelectedBotChange :one
INSERT INTO twitch.selected_bots (user_id, bot_id, broadcaster_id)
VALUES ($1, $2, $3)
ON CONFLICT (user_id) DO UPDATE
  SET 
  bot_id = $2,
  broadcaster_id = $3,
  updated_at = CURRENT_TIMESTAMP
  RETURNING user_id, broadcaster_id, bot_id, updated_at
`

type TwitchSelectedBotChangeParams struct {
	UserID        uuid.UUID
	BotID         string
	BroadcasterID string
}

func (q *Queries) TwitchSelectedBotChange(ctx context.Context, arg TwitchSelectedBotChangeParams) (TwitchSelectedBot, error) {
	row := q.db.QueryRow(ctx, twitchSelectedBotChange, arg.UserID, arg.BotID, arg.BroadcasterID)
	var i TwitchSelectedBot
	err := row.Scan(
		&i.UserID,
		&i.BroadcasterID,
		&i.BotID,
		&i.UpdatedAt,
	)
	return i, err
}

const twitchSelectedBotGetByBroadcasterID = `-- name: TwitchSelectedBotGetByBroadcasterID :one
SELECT user_id, broadcaster_id, bot_id, updated_at
FROM twitch.selected_bots
WHERE broadcaster_id = $1
`

func (q *Queries) TwitchSelectedBotGetByBroadcasterID(ctx context.Context, broadcasterID string) (TwitchSelectedBot, error) {
	row := q.db.QueryRow(ctx, twitchSelectedBotGetByBroadcasterID, broadcasterID)
	var i TwitchSelectedBot
	err := row.Scan(
		&i.UserID,
		&i.BroadcasterID,
		&i.BotID,
		&i.UpdatedAt,
	)
	return i, err
}

const twitchSelectedBotGetByUserID = `-- name: TwitchSelectedBotGetByUserID :one
SELECT user_id, broadcaster_id, bot_id, updated_at
FROM twitch.selected_bots
WHERE user_id = $1
`

func (q *Queries) TwitchSelectedBotGetByUserID(ctx context.Context, userID uuid.UUID) (TwitchSelectedBot, error) {
	row := q.db.QueryRow(ctx, twitchSelectedBotGetByUserID, userID)
	var i TwitchSelectedBot
	err := row.Scan(
		&i.UserID,
		&i.BroadcasterID,
		&i.BotID,
		&i.UpdatedAt,
	)
	return i, err
}

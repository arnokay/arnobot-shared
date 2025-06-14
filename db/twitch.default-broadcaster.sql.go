// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: twitch.default-broadcaster.sql

package db

import (
	"context"
)

const twitchDefaultBroadcasterGet = `-- name: TwitchDefaultBroadcasterGet :one
SELECT main, twitch_user_id
FROM twitch.default_broadcaster
WHERE main = true
`

func (q *Queries) TwitchDefaultBroadcasterGet(ctx context.Context) (TwitchDefaultBroadcaster, error) {
	row := q.db.QueryRow(ctx, twitchDefaultBroadcasterGet)
	var i TwitchDefaultBroadcaster
	err := row.Scan(&i.Main, &i.TwitchUserID)
	return i, err
}

const twitchDefaultBroadcasterUpdate = `-- name: TwitchDefaultBroadcasterUpdate :execrows
INSERT INTO twitch.default_broadcaster (main, twitch_user_id)
VALUES (true, $1)
ON CONFLICT (main) DO UPDATE
  SET twitch_user_id = $1
`

func (q *Queries) TwitchDefaultBroadcasterUpdate(ctx context.Context, twitchUserID string) (int64, error) {
	result, err := q.db.Exec(ctx, twitchDefaultBroadcasterUpdate, twitchUserID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

package data

import (
	"time"

	"arnobot-shared/db"
)

type TwitchDefaultBot struct {
	TwitchUserID string
}

func NewTwitchDefaultBotFromDB(fromDB db.TwitchDefaultBot) TwitchDefaultBot {
	return TwitchDefaultBot{
		TwitchUserID: fromDB.TwitchUserID,
	}
}

type TwitchSelectedBot struct {
	UserID       int32
	TwitchUserID string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewTwitchSelectedBotFromDB(fromDB db.TwitchSelectedBot) TwitchSelectedBot {
	return TwitchSelectedBot{
		UserID:       fromDB.UserID,
		TwitchUserID: fromDB.TwitchUserID,
		CreatedAt:    fromDB.CreatedAt,
		UpdatedAt:    fromDB.UpdatedAt,
	}
}

type TwitchBotRole string

const (
	TwitchBotRoleUser        TwitchBotRole = "user"
	TwitchBotRoleVIP         TwitchBotRole = "vip"
	TwitchBotRoleModerator   TwitchBotRole = "moderator"
	TwitchBotRoleBroadcaster TwitchBotRole = "broadcaster"
)

type TwitchBot struct {
	UserID       int32
	TwitchUserID string
	Role         TwitchBotRole
}

func NewTwitchBotFromDB(fromDB db.TwitchBot) TwitchBot {
	return TwitchBot{
		UserID:       fromDB.UserID,
		TwitchUserID: fromDB.TwitchUserID,
		Role:         TwitchBotRole(fromDB.Role),
	}
}

package data

import (
	"arnobot-shared/db"
	"time"
)

type TwitchDefaultBot struct {
	TwitchUserID string
}

type TwitchSelectedBot struct {
	UserID       int32
	TwitchUserID string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewTwitchDefaultBotFromDB(fromDB db.TwitchDefaultBot) TwitchDefaultBot {
	return TwitchDefaultBot{
		TwitchUserID: fromDB.TwitchUserID,
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

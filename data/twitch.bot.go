package data

import "arnobot-shared/db"

type TwitchDefaultBot struct {
	TwitchUserID string
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
	UserID       int
	TwitchUserID string
	Role         TwitchBotRole
}

func NewTwitchBotFromDB(fromDB db.TwitchBot) TwitchBot {
	return TwitchBot{
		UserID:       int(fromDB.UserID),
		TwitchUserID: fromDB.TwitchUserID,
		Role:         TwitchBotRole(fromDB.Role),
	}
}

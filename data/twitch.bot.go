package data

import (
	"time"

	"github.com/google/uuid"

	"arnobot-shared/db"
)

type TwitchDefaultBot struct {
	BotID string
}

func NewTwitchDefaultBotFromDB(fromDB db.TwitchDefaultBot) TwitchDefaultBot {
	return TwitchDefaultBot{
		BotID: fromDB.BotID,
	}
}

type TwitchSelectedBot struct {
	UserID        uuid.UUID
	BotID         string
	BroadcasterID string
	UpdatedAt     time.Time
}

func NewTwitchSelectedBotFromDB(fromDB db.TwitchSelectedBot) TwitchSelectedBot {
	return TwitchSelectedBot{
		UserID:        fromDB.UserID,
		BotID:         fromDB.BotID,
		BroadcasterID: fromDB.BroadcasterID,
		UpdatedAt:     fromDB.UpdatedAt,
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
	UserID        uuid.UUID
	BotID         string
	BroadcasterID string
	Role          TwitchBotRole
}

func NewTwitchBotFromDB(fromDB db.TwitchBot) TwitchBot {
	return TwitchBot{
		UserID:        fromDB.UserID,
		BroadcasterID: fromDB.BroadcasterID,
		BotID:         fromDB.BotID,
		Role:          TwitchBotRole(fromDB.Role),
	}
}

type TwitchBotCreate struct {
	UserID        uuid.UUID
	BotID         string
	BroadcasterID string
}

func (d TwitchBotCreate) ToDB() db.TwitchBotCreateParams {
	return db.TwitchBotCreateParams{
		UserID:        d.UserID,
		BotID:         d.BotID,
		BroadcasterID: d.BroadcasterID,
	}
}

type TwitchBotsGet struct {
	UserID        *uuid.UUID
	BotID         *string
	BroadcasterID *string
}

func (d TwitchBotsGet) ToDB() db.TwitchBotsGetParams {
	return db.TwitchBotsGetParams{
		UserID:        d.UserID,
		BotID:         d.BotID,
		BroadcasterID: d.BroadcasterID,
	}
}

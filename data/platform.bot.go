package data

import (
	"time"

	"github.com/google/uuid"

	"github.com/arnokay/arnobot-shared/platform"
)

type PlatformBotToggle struct {
	Platform platform.Platform `json:"platform" param:"platform" validate:"validateFn"`
	UserID   uuid.UUID         `json:"userId"`
}

type PlatformBotGet struct {
	Platform platform.Platform
	UserID   uuid.UUID
}

type PlatformBot struct {
	Platform      platform.Platform
	BotID         string
	BroadcasterID string
	UserID        uuid.UUID
}

type PlatformDefaultBot struct {
	BotID string
}

type PlatformSelectedBot struct {
	UserID        uuid.UUID
	BotID         string
	BroadcasterID string
	Enabled       bool
	UpdatedAt     time.Time
}

type PlatformBotCreate struct {
	UserID        uuid.UUID
	BotID         string
	BroadcasterID string
}

type PlatformBotsGet struct {
	UserID        *uuid.UUID
	BotID         *string
	BroadcasterID *string
}

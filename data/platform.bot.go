package data

import (
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
	Platform platform.Platform
	BotID    string
	UserID   uuid.UUID
	Enabled  bool
}

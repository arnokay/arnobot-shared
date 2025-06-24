package data

import (
	"github.com/google/uuid"

	"github.com/arnokay/arnobot-shared/platform"
)

type PlatformToggleBot struct {
	Platform platform.Platform
	UserID   uuid.UUID
}

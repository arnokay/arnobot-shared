package data

import (
	"github.com/google/uuid"

	"github.com/arnokay/arnobot-shared/platform"
)

type PlatformToggleBot struct {
	Platform platform.Platform `json:"platform" param:"platform" validate:"validateFn=IsEnum"`
	UserID   uuid.UUID         `json:"userId" validate:"required,uuid4"`
}

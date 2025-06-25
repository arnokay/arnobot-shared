package data

import (
	"github.com/google/uuid"

	"github.com/arnokay/arnobot-shared/platform"
)

type PlatformUser struct {
	ID       string
	Name     string
	Login    string
	UserID   *uuid.UUID
	Platform platform.Platform
}

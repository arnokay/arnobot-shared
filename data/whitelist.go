package data

import (
	"github.com/google/uuid"

	"github.com/arnokay/arnobot-shared/db"
	"github.com/arnokay/arnobot-shared/platform"
)

type Whitelist struct {
	Platform          platform.Platform `json:"platform"`
	PlatformUserID    *string           `json:"platformUserId"`
	PlatformUserName  *string           `json:"platformUserName"`
	PlatformUserLogin *string           `json:"platformUserLogin"`
	UserID            *uuid.UUID        `json:"userId"`
}

func NewWhitelistFromDB(fromDB db.Whitelist) Whitelist {
	return Whitelist{
		Platform:          platform.Platform(fromDB.Platform),
		PlatformUserID:    fromDB.PlatformUserID,
		PlatformUserName:  fromDB.PlatformUserName,
		PlatformUserLogin: fromDB.PlatformUserLogin,
		UserID:            fromDB.UserID,
	}
}

type WhitelistCreate struct {
	Platform          string     `json:"platform"`
	PlatformUserID    *string    `json:"platformUserId"`
	PlatformUserName  *string    `json:"platformUserName"`
	PlatformUserLogin *string    `json:"platformUserLogin"`
	UserID            *uuid.UUID `json:"userId"`
}

type WhitelistGetOne struct {
	Platform          string     `json:"platform"`
	PlatformUserID    *string    `json:"platformUserId"`
	PlatformUserName  *string    `json:"platformUserName"`
	PlatformUserLogin *string    `json:"platformUserLogin"`
	UserID            *uuid.UUID `json:"userId"`
}

func (w WhitelistGetOne) ToDB() db.WhitelistGetOneParams {
	return db.WhitelistGetOneParams{
		Platform:          db.Platform(w.Platform),
		PlatformUserID:    w.PlatformUserID,
		PlatformUserName:  w.PlatformUserName,
		PlatformUserLogin: w.PlatformUserLogin,
		UserID:            w.UserID,
	}
}

type WhitelistDelete struct {
	Platform          string     `json:"platform"`
	PlatformUserID    *string    `json:"platformUserId"`
	PlatformUserName  *string    `json:"platformUserName"`
	PlatformUserLogin *string    `json:"platformUserLogin"`
	UserID            *uuid.UUID `json:"userId"`
}

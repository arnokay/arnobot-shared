package data

import (
	"time"

	"arnobot-shared/db"
)

type TwitchUser struct {
	ID              string
	Username        string
	DisplayName     string
	Type            string
	BroadcasterType string
	ProfileImageUrl string
	CreatedAt       time.Time
	AuthProviderID  *int
}

type TwitchUserCreate struct {
	ID              string
	Username        string
	DisplayName     string
	Type            string
	BroadcasterType string
	ProfileImageUrl string
	CreatedAt       time.Time
}

func (u TwitchUserCreate) ToDB() db.TwitchUserCreateParams {
	return db.TwitchUserCreateParams{
		ID:              u.ID,
		Username:        u.Username,
		DisplayName:     u.DisplayName,
		Type:            u.Type,
		BroadcasterType: u.BroadcasterType,
		ProfileImageUrl: u.ProfileImageUrl,
		CreatedAt:       u.CreatedAt,
	}
}

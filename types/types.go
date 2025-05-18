package types

import (
	"arnobot-shared/data"
)

type FullUserPayload struct {
	User       data.User
	TwitchUser data.TwitchUser
}

type AccessRefreshTokenPayload struct {
	Provider data.AuthProvider
}

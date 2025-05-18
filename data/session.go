package data

import (
	"time"

	"arnobot-shared/db"
)

type AuthSession struct {
	Token      string    `json:"token"`
	UserID     int       `json:"userId"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	LastUsedAt time.Time `json:"lastUsedAt"`
}

func NewSessionFromDB(fromDB db.AuthSession) AuthSession {
	session := AuthSession{
		Token:      fromDB.Token,
		UserID:     int(fromDB.UserID),
		Status:     string(fromDB.Status),
		CreatedAt:  fromDB.CreatedAt.Time,
		LastUsedAt: fromDB.CreatedAt.Time,
	}

	return session
}

package data

import (
	"time"

	"arnobot-shared/db"
)

type AuthSessionStatus string

const (
	AuthSessionStatusActive   AuthSessionStatus = "active"
	AuthSessionStatusDisabled AuthSessionStatus = "disabled"
)

type AuthSession struct {
	Token      string            `json:"token"`
	UserID     int32             `json:"userId"`
	Status     AuthSessionStatus `json:"status"`
	CreatedAt  time.Time         `json:"createdAt"`
	LastUsedAt time.Time         `json:"lastUsedAt"`
}

func NewSessionFromDB(fromDB db.AuthSession) AuthSession {
	session := AuthSession{
		Token:      fromDB.Token,
		UserID:     fromDB.UserID,
		Status:     AuthSessionStatus(fromDB.Status),
		CreatedAt:  fromDB.CreatedAt,
		LastUsedAt: fromDB.CreatedAt,
	}

	return session
}

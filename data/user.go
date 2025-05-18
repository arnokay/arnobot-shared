package data

import (
	"arnobot-shared/db"
	"context"
	"time"
)

type User struct {
  ID        int       `json:"id"`
	Username  string    `json:"username"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUserFromDB(fromDB db.User) User {
  return User{
    ID: int(fromDB.ID),
    Username: fromDB.Username,
    Status: string(fromDB.Status),
    CreatedAt: fromDB.CreatedAt.Time,
    UpdatedAt: fromDB.UpdatedAt.Time,
  }
} 

type UserCreate struct {
	Username string `json:"username"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Status   string `json:"status"`
}

type Command struct {
	Cmd             string
	Aliases         []string
	Description     string
	DescriptionFunc func(context.Context) string
}

type TwitchEvent struct {
	ID          string
	Username    string
	DisplayName string
	Channel     string
	Msg         string
}

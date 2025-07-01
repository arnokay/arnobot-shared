package data

import (
	"time"

	"github.com/google/uuid"

	"github.com/arnokay/arnobot-shared/db"
)

type UserCommand struct {
	UserID    uuid.UUID `json:"userId"`
	Name      string    `json:"name"`
	Text      string    `json:"text"`
	Reply     bool      `json:"reply"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUserCommandFromDB(fromDB db.CoreUserCommand) UserCommand {
	return UserCommand{
		UserID:    fromDB.UserID,
		Name:      fromDB.Name,
		Text:      fromDB.Text,
		Reply:     fromDB.Reply,
		CreatedAt: fromDB.CreatedAt,
		UpdatedAt: fromDB.UpdatedAt,
	}
}

type UserCommandGetOne struct {
	UserID uuid.UUID `json:"userId"`
	Name   string    `json:"name"`
}

type UserCommandCreate struct {
	UserID uuid.UUID `json:"userId"`
	Name   string    `json:"name"`
	Text   string    `json:"text"`
	Reply  bool      `json:"reply"`
}

type UserCommandUpdate struct {
	UserID  uuid.UUID `json:"userId"`
	Name    string    `json:"name"`
	NewName *string   `json:"newName"`
	Text    *string   `json:"text"`
	Reply   *bool     `json:"reply"`
}

type UserCommandDelete struct {
	UserID uuid.UUID `json:"userId"`
	Name   string    `json:"name"`
}

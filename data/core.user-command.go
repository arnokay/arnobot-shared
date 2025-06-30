package data

import (
	"github.com/google/uuid"

	"github.com/arnokay/arnobot-shared/db"
)

type UserCommand struct {
	UserID uuid.UUID `json:"userId"`
	Name   string    `json:"name"`
	Text   string    `json:"text"`
	Reply  bool      `json:"reply"`
}

func NewUserCommandFromDB(fromDB db.CoreUserCommand) UserCommand {
	return UserCommand{
		UserID: fromDB.UserID,
		Name:   fromDB.Name,
		Text:   fromDB.Text,
		Reply:  fromDB.Reply,
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
	UserID  uuid.UUID  `json:"userId"`
	Name    string     `json:"name"`
	NewName *uuid.UUID `json:"newName"`
	Text    *string    `json:"text"`
	Reply   *bool      `json:"reply"`
}

type UserCommandDelete struct {
	UserID uuid.UUID `json:"userId"`
	Name   string    `json:"name"`
}

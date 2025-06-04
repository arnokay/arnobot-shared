package data

import (
	"time"

	"github.com/google/uuid"

	"arnobot-shared/db"
)

type UserStatus = db.UserStatus

type User struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Status    UserStatus `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func NewUserFromDB(fromDB db.User) User {
	return User{
		ID:        fromDB.ID,
		Username:  fromDB.Username,
		Status:    fromDB.Status,
		CreatedAt: fromDB.CreatedAt,
		UpdatedAt: fromDB.UpdatedAt,
	}
}

type UserUpdate struct {
	Username *string     `json:"username"`
	Status   *UserStatus `json:"status"`
}

func (u UserUpdate) ToDB(id uuid.UUID) db.UserUpdateParams {
	status := db.NullUserStatus{
		Valid: false,
	}

	if u.Status != nil {
		status.UserStatus = *u.Status
		status.Valid = true
	}

	return db.UserUpdateParams{
		ID:       id,
		Username: u.Username,
		Status:   status,
	}
}

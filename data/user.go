package data

import (
	"time"

	"arnobot-shared/db"
)

type UserStatus = db.UserStatus

type User struct {
	ID        int32      `json:"id"`
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

func (u UserUpdate) ToDB(id int32) db.UserUpdateParams {
  status := db.NullUserStatus{
    Valid: false,
  }

  if u.Status != nil {
    status.UserStatus = *u.Status
    status.Valid = true
  }

  return db.UserUpdateParams{
    ID: id,
    Username: u.Username,
    Status: status,
  }
}

package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	ADMIN_ROLE  Role = "admin"
	CLIENT_ROLE Role = "client"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"password" gorm:"not null"`
	Role     Role      `json:"role" gorm:"type:varchar(50);not null"`

	Phone   string `json:"phone" gorm:"type:varchar(20)"`
	Gender  string `json:"gender" gorm:"type:varchar(20)"`
	Address string `json:"address" gorm:"type:text"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func IsValidRole(r Role) bool {
	switch r {
	case ADMIN_ROLE, CLIENT_ROLE:
		return true
	}
	return false
}

func UserIsAdmin(u *User) error {
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return fmt.Errorf("missing requireds fields for admin")
	}
	return nil
}

func UserIsClient(u *User) error {
	if u.Name == "" ||
		u.Email == "" ||
		u.Password == "" ||
		u.Phone == "" ||
		u.Gender == "" ||
		u.Address == "" {
		return fmt.Errorf("all client fields are required")
	}
	return nil
}

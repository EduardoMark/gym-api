package equipament

import (
	"time"

	"github.com/google/uuid"
)

type Category string

const (
	CATEGORY_CARDIO   Category = "cardio"
	CATEGORY_STRENGTH Category = "strength"
)

type Status string

const (
	STATUS_ACTIVE      Status = "active"
	STATUS_INACTIVE    Status = "inactive"
	STATUS_MAINTENANCE Status = "in_maintenance"
)

type Equipament struct {
	ID              uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name            string     `json:"name" gorm:"not null"`
	Description     *string    `json:"description"`
	Category        Category   `json:"category" gorm:"not null"`
	Brand           string     `json:"brand" gorm:"not null"`
	Model           string     `json:"model" gorm:"not null"`
	MaintenanceDate *time.Time `json:"maintenance_date"`
	Status          Status     `json:"status" gorm:"not null"`
	Quantity        int        `json:"quantity" gorm:"not null;default:1"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func IsValidCategory(c Category) bool {
	switch c {
	case CATEGORY_CARDIO, CATEGORY_STRENGTH:
		return true
	}
	return false
}

func IsValidStatus(s Status) bool {
	switch s {
	case STATUS_ACTIVE, STATUS_INACTIVE, STATUS_MAINTENANCE:
		return true
	}
	return false
}

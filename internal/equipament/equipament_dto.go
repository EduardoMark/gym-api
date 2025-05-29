package equipament

import (
	"time"

	"github.com/google/uuid"
)

type EquipamentRequest struct {
	Name            string     `json:"name"`
	Description     *string    `json:"description"`
	Category        Category   `json:"category"`
	Brand           string     `json:"brand"`
	Model           string     `json:"model"`
	MaintenanceDate *time.Time `json:"maintenance_date"`
	Status          Status     `json:"status"`
	Quantity        int        `json:"quantidade"`
}

type EquipamentResponse struct {
	ID              uuid.UUID  `json:"id"`
	Name            string     `json:"name"`
	Description     *string    `json:"description"`
	Category        Category   `json:"category"`
	Brand           string     `json:"brand"`
	Model           string     `json:"model"`
	MaintenanceDate *time.Time `json:"maintenance_date"`
	Status          Status     `json:"status"`
	Quantity        int        `json:"quantidade"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

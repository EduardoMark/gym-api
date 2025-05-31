package plan

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PlanRequest struct {
	Name             string          `json:"name"`
	Description      *string         `json:"description"`
	Price            decimal.Decimal `json:"price"`
	Cicle            Cicle           `json:"cicle"`
	DurationInMonths int             `json:"duration_in_months"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

type PlanResponse struct {
	ID               uuid.UUID       `json:"id"`
	Name             string          `json:"name"`
	Description      *string         `json:"description"`
	Price            decimal.Decimal `json:"price"`
	Cicle            Cicle           `json:"cicle"`
	DurationInMonths int             `json:"duration_in_months"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

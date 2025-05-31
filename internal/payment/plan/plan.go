package plan

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Cicle string

const (
	CicleWeekly       Cicle = "weekly"
	CicleBiweekly     Cicle = "biweekly"
	CicleMonthly      Cicle = "monthly"
	CicleBimonthly    Cicle = "bimonthly"
	CicleQuarterly    Cicle = "quarterly"
	CicleSemiannually Cicle = "semiannually"
	CicleYearly       Cicle = "yearly"
)

type Plan struct {
	ID               uuid.UUID       `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name             string          `gorm:"not null"`
	Description      *string         `gorm:"type:text;"`
	Price            decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	Cicle            Cicle           `gorm:"not null"`
	DurationInMonths int             `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func IsValidCycle(c Cicle) bool {
	switch c {
	case CicleWeekly, CicleBiweekly, CicleMonthly, CicleBimonthly,
		CicleQuarterly, CicleSemiannually, CicleYearly:
		return true
	}
	return false
}

package plan

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type UseCase interface {
	Create(params *PlanRequest) error
	FindOne(id string) (*Plan, error)
	FindAll() ([]Plan, error)
	Update(id string, params *PlanRequest) error
	Delete(id string) error
}

type planUseCase struct {
	repo Repository
}

func NewPlanUseCase(repo Repository) UseCase {
	return &planUseCase{repo: repo}
}

func (uc *planUseCase) Create(params *PlanRequest) error {
	if !IsValidCycle(params.Cicle) {
		return fmt.Errorf("invalid type cicle")
	}
	if params.Price.LessThan(decimal.Zero) {
		return fmt.Errorf("the price cannot be less than 0")
	}
	if params.DurationInMonths <= 0 {
		return fmt.Errorf("the duration in months cannot be less than 1")
	}

	plan := Plan{
		ID:               uuid.New(),
		Name:             params.Name,
		Description:      params.Description,
		Price:            params.Price,
		Cicle:            params.Cicle,
		DurationInMonths: params.DurationInMonths,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := uc.repo.Create(&plan); err != nil {
		return err
	}

	return nil
}

func (uc *planUseCase) FindOne(id string) (*Plan, error) {
	plan, err := uc.repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (uc *planUseCase) FindAll() ([]Plan, error) {
	plans, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return plans, nil
}

func (uc *planUseCase) Update(id string, params *PlanRequest) error {
	plan, err := uc.repo.FindOne(id)
	if err != nil {
		return err
	}

	if params.Name != "" {
		plan.Name = params.Name
	}
	if params.Description != nil {
		plan.Description = params.Description
	}
	if params.Price != decimal.Zero && params.Price.GreaterThanOrEqual(decimal.Zero) {
		plan.Price = params.Price
	}
	if params.Cicle != "" && IsValidCycle(params.Cicle) {
		plan.Cicle = params.Cicle
	}
	if params.DurationInMonths > 0 {
		plan.DurationInMonths = params.DurationInMonths
	}

	if err := uc.repo.Update(id, plan); err != nil {
		return err
	}

	return nil
}

func (uc *planUseCase) Delete(id string) error {
	if err := uc.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

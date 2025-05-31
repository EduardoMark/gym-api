package equipament

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UseCase interface {
	Create(params EquipamentRequest) error
	FindOne(id string) (*Equipament, error)
	FindAll() ([]Equipament, error)
	Update(id string, params EquipamentRequest) error
	Delete(id string) error
}

type equipamentUseCase struct {
	repo Repository
}

func NewEquipamentUseCase(repo Repository) UseCase {
	return &equipamentUseCase{repo: repo}
}

func (uc *equipamentUseCase) Create(params EquipamentRequest) error {
	if !IsValidCategory(params.Category) {
		return fmt.Errorf("invalid category")
	}

	if !IsValidStatus(params.Status) {
		return fmt.Errorf("invalid status")
	}

	equipament := Equipament{
		ID:              uuid.New(),
		Name:            params.Name,
		Description:     params.Description,
		Category:        params.Category,
		Brand:           params.Brand,
		Model:           params.Model,
		MaintenanceDate: params.MaintenanceDate,
		Status:          params.Status,
		Quantity:        params.Quantity,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := uc.repo.Create(&equipament); err != nil {
		return err
	}

	return nil
}

func (uc *equipamentUseCase) FindOne(id string) (*Equipament, error) {
	equipament, err := uc.repo.FindOne(id)
	if err != nil {
		return nil, err
	}
	return equipament, nil
}

func (uc *equipamentUseCase) FindAll() ([]Equipament, error) {
	equipaments, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return equipaments, nil
}

func (uc *equipamentUseCase) Update(id string, params EquipamentRequest) error {
	equipament, err := uc.repo.FindOne(id)
	if err != nil {
		return err
	}

	if params.Name != "" {
		equipament.Name = params.Name
	}
	if params.Description != nil {
		equipament.Description = params.Description
	}
	if params.Brand != "" {
		equipament.Brand = params.Brand
	}
	if params.Category != "" && IsValidCategory(params.Category) {
		equipament.Category = params.Category
	}
	if params.Status != "" && IsValidStatus(params.Status) {
		equipament.Status = params.Status
	}
	if params.MaintenanceDate != nil {
		equipament.MaintenanceDate = params.MaintenanceDate
	}
	if params.Model != "" {
		equipament.Model = params.Model
	}
	if params.Quantity != 0 {
		equipament.Quantity = params.Quantity
	}

	if err := uc.repo.Update(equipament); err != nil {
		return err
	}

	return nil
}

func (uc *equipamentUseCase) Delete(id string) error {
	if err := uc.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

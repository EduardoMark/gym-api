package equipament

import "fmt"

type UseCase interface {
	Create(params EquipamentRequest) error
	FindOne(id string) (*Equipament, error)
	FindAll() ([]Equipament, error)
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
		Name:            params.Name,
		Description:     params.Description,
		Category:        params.Category,
		Brand:           params.Brand,
		Model:           params.Model,
		MaintenanceDate: params.MaintenanceDate,
		Status:          params.Status,
		Quantity:        params.Quantity,
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

func (uc *equipamentUseCase) Delete(id string) error {
	if err := uc.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

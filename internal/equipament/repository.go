package equipament

import "gorm.io/gorm"

type Repository interface {
	Create(equipament *Equipament) error
	FindOne(id string) (*Equipament, error)
	FindAll() ([]Equipament, error)
	Update(equipament *Equipament) error
	Delete(id string) error
}

type equipamentRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &equipamentRepository{db: db}
}

func (r *equipamentRepository) Create(equipament *Equipament) error {
	if err := r.db.Create(equipament).Error; err != nil {
		return err
	}
	return nil
}

func (r *equipamentRepository) FindOne(id string) (*Equipament, error) {
	var equipament *Equipament

	if err := r.db.Where("id = ?", id).First(&equipament).Error; err != nil {
		return nil, err
	}
	return equipament, nil
}

func (r *equipamentRepository) FindAll() ([]Equipament, error) {
	var equipaments []Equipament

	if err := r.db.Find(&equipaments).Error; err != nil {
		return nil, err
	}
	return equipaments, nil
}

func (r *equipamentRepository) Update(equipament *Equipament) error {
	if err := r.db.Updates(equipament).Error; err != nil {
		return err
	}
	return nil
}

func (r *equipamentRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&Equipament{}).Error; err != nil {
		return err
	}
	return nil
}

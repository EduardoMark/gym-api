package plan

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(plan *Plan) error
	FindOne(id string) (*Plan, error)
	FindAll() ([]Plan, error)
	Update(id string, plan *Plan) error
	Delete(id string) error
}

type planRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &planRepository{db: db}
}

func (r *planRepository) Create(plan *Plan) error {
	if err := r.db.Create(plan).Error; err != nil {
		return err
	}
	return nil
}

func (r *planRepository) FindOne(id string) (*Plan, error) {
	var plan *Plan

	if err := r.db.Where("id = ?", id).First(&plan).Error; err != nil {
		return nil, err
	}

	return plan, nil
}

func (r *planRepository) FindAll() ([]Plan, error) {
	var plans []Plan

	if err := r.db.Find(&plans).Error; err != nil {
		return nil, err
	}

	return plans, nil
}

func (r *planRepository) Update(id string, plan *Plan) error {
	plan.UpdatedAt = time.Now()
	if err := r.db.Updates(plan).Error; err != nil {
		return err
	}
	return nil
}

func (r *planRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&Plan{}).Error; err != nil {
		return err
	}
	return nil
}

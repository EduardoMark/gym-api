package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

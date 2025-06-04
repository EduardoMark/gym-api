package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	FindOne(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll() ([]User, error)
	Update(user *User) error
	Delete(id string) error
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

func (r *userRepository) FindOne(id string) (*User, error) {
	var user *User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*User, error) {
	var user *User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindAll() ([]User, error) {
	var user []User

	if err := r.db.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(user *User) error {
	if err := r.db.Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(id string) error {
	if err := r.db.Where("id=?", id).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

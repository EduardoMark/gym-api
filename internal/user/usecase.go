package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	Create(params UserRequest) error
	FindOne(id string) (*User, error)
	FindAll() ([]User, error)
	Update(id string, params UserRequest) error
	Delete(id string) error
}

type userUseCase struct {
	repo Repository
}

func NewUserUseCase(repo Repository) UseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (uc *userUseCase) Create(params UserRequest) error {
	if !IsValidRole(params.Role) {
		return fmt.Errorf("invalid role")
	}

	userID := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		return fmt.Errorf("error on hashing password: %w", err)
	}

	user := User{
		ID:        userID,
		Name:      params.Name,
		Email:     params.Email,
		Password:  string(hashedPassword),
		Role:      params.Role,
		Phone:     params.Phone,
		Gender:    params.Gender,
		Address:   params.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	switch user.Role {
	case ADMIN_ROLE:
		if err := UserIsAdmin(&user); err != nil {
			return err
		}
	case CLIENT_ROLE:
		if err := UserIsClient(&user); err != nil {
			return err
		}
	}

	return uc.repo.Create(&user)
}

func (uc *userUseCase) FindOne(id string) (*User, error) {
	user, err := uc.repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) FindAll() ([]User, error) {
	user, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) Update(id string, params UserRequest) error {
	return nil
}

func (uc *userUseCase) Delete(id string) error {
	if err := uc.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

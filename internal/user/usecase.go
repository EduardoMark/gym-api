package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	Create(params CreateUserParams) error
}

type userUseCase struct {
	repo Repository
}

func NewUserUseCase(repo Repository) UseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (uc *userUseCase) Create(params CreateUserParams) error {
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

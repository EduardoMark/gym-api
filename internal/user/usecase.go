package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/EduardoMark/gym-api/pkg/auth"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	Login(params UserLoginRequest) (string, error)
	Create(params UserRequest) error
	FindOne(id string) (*User, error)
	FindAll() ([]User, error)
	Update(id string, params UserRequest) error
	Delete(id string) error
}

type userUseCase struct {
	repo Repository
	auth auth.Authorizer
}

func NewUserUseCase(repo Repository, auth auth.Authorizer) UseCase {
	return &userUseCase{
		repo: repo,
		auth: auth,
	}
}

func (uc *userUseCase) Login(params UserLoginRequest) (string, error) {
	if params.Email == "" || params.Password == "" {
		return "", fmt.Errorf("all fields is required")
	}

	record, err := uc.repo.FindByEmail(params.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	isValid := bcrypt.CompareHashAndPassword([]byte(record.Password), []byte(params.Password))
	if isValid != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := uc.auth.GenerateToken(record.Name, string(record.Role))
	if err != nil {
		return "", errors.New("error on generating token")
	}

	return token, nil
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
	user, err := uc.repo.FindOne(id)
	if err != nil {
		return err
	}

	if !IsValidRole(params.Role) {
		return fmt.Errorf("invalid role")
	}
	if params.Role != "" && IsValidRole(params.Role) {
		user.Role = params.Role
	}

	if params.Name != "" {
		user.Name = params.Name
	}
	if params.Email != "" {
		user.Email = params.Email
	}
	if params.Address != "" {
		user.Address = params.Address
	}
	if params.Gender != "" {
		user.Gender = params.Gender
	}
	if params.Phone != "" {
		user.Phone = params.Phone
	}

	if err := UserIsClient(user); err != nil {
		return err
	}

	if err := uc.repo.Update(user); err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) Delete(id string) error {
	if err := uc.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

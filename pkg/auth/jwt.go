package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Authorizer interface {
	GenerateToken(name, role string) (string, error)
	VerifyToken(tokenString string) error
}

type JWTAuth struct {
	SecretKey string
}

func NewJWTAuth(sk string) Authorizer {
	return &JWTAuth{SecretKey: sk}
}

func (j *JWTAuth) GenerateToken(name, role string) (string, error) {
	secretKey := []byte(j.SecretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWTAuth) VerifyToken(tokenString string) error {
	secretKey := []byte(j.SecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

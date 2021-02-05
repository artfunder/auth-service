package service

import (
	"github.com/artfunder/structs"
)

// AuthService ...
type AuthService struct {
	users UserGetter
}

// NewAuthService ...
func NewAuthService(userGetter UserGetter) *AuthService {
	s := new(AuthService)
	s.users = userGetter
	return s
}

// UserGetter ...
type UserGetter interface {
	GetByUsername(username string) (structs.User, error)
	GetByEmail(email string) (structs.User, error)
}

// LocalLogin takes a username/email and a password, and returns a token or error
func (s AuthService) LocalLogin(usernameOrEmail string, password string) (string, error) {
	if !s.isUsernameValid(usernameOrEmail) && !s.isEmailValid(usernameOrEmail) {
		return "", ErrUserNotFound
	}

	return "token", nil
}

func (s AuthService) isUsernameValid(username string) bool {
	_, err := s.users.GetByUsername(username)
	if err != nil {
		return false
	}
	return true
}

func (s AuthService) isEmailValid(email string) bool {
	_, err := s.users.GetByEmail(email)
	if err != nil {
		return false
	}
	return true
}

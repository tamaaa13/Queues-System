package admin

import (
	"errors"

	"github.com/tamaaa13/Queue/internal/model/admin"
	"github.com/tamaaa13/Queue/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(username, password string) (string, *admin.User, error) {
	user, err := s.adminRepo.FindUserByUsername(username)
	if err != nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, errors.New("Username atau Password Salah !")
	}

	token, err := jwt.GenerateJWT(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

// HashPassword untuk menyimpan password dalam bentuk hash
func (s *service) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

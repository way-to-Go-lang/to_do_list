package service

import (
	"crypto/sha1"
	"fmt"
	todo "to_do_list"
	"to_do_list/pkg/repository"
)

const salt = "hadfshehasfehihf3442h"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprint("%x", hash.Sum([]byte(salt)))
}

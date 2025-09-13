package service

import (
	todo "to_do_list"
	"to_do_list/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type Todolist interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	Todolist
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

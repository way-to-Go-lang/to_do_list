package repository

import (
	"github.com/jmoiron/sqlx"
	todo "to_do_list"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type Todolist interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	Todolist
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

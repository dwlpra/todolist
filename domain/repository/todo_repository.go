package repository

import "github.com/dwlpra/todolist/domain/entity"

type TodoRepository interface {
	FindAll(*string) (*[]entity.Todo, error)
	FindByID(int) (*entity.Todo, error)
	Save(*entity.Todo) error
	Update(int, *entity.Todo) error
	Delete(int) error
}

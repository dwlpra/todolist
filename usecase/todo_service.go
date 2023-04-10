package usecase

import (
	"github.com/dwlpra/todolist/domain/entity"
	"github.com/dwlpra/todolist/domain/repository"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo}
}

func (s *TodoService) GetAllTodos(activityGroupID *string) (*[]entity.Todo, error) {
	return s.repo.FindAll(activityGroupID)
}

func (s *TodoService) GetTodoByID(id int) (*entity.Todo, error) {
	return s.repo.FindByID(id)
}

func (s *TodoService) CreateTodo(Todo *entity.Todo) error {
	return s.repo.Save(Todo)
}

func (s *TodoService) UpdateTodo(id int, todo *entity.Todo) (*entity.Todo, error) {
	s.repo.Update(id, todo)
	return s.repo.FindByID(id)
}

func (s *TodoService) DeleteTodo(id int) error {
	return s.repo.Delete(id)
}

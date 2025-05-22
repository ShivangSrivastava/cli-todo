package usecase

import (
	"errors"
	"strings"

	"github.com/ShivangSrivastava/cli-todo/internal/entity"
	"github.com/ShivangSrivastava/cli-todo/internal/interfaces"
)

type TodoUsecase struct {
	repo interfaces.TodoRepository
}

func NewTodoUsecase(repo interfaces.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}

func (u *TodoUsecase) GetTodos() ([]entity.Todo, error) {
	return u.repo.GetTodos()
}

func (u *TodoUsecase) CreateTodo(title string) (entity.Todo, error) {
	trimmedTitle := strings.TrimSpace(title)
	if len(trimmedTitle) <= 5 {
		return entity.Todo{}, errors.New("title should be more than 5 characters")
	}
	return u.repo.CreateTodo(trimmedTitle)
}

func (u *TodoUsecase) UpdateTodoTitle(id int, title string) (entity.Todo, error) {
	trimmedTitle := strings.TrimSpace(title)
	if len(trimmedTitle) <= 5 {
		return entity.Todo{}, errors.New("title should be more than 5 characters")
	}
	return u.repo.UpdateTodoTitle(id, title)
}

func (u *TodoUsecase) UpdateTodoIsCompleted(id int, isCompleted bool) (entity.Todo, error) {
	return u.repo.UpdateTodoIsCompleted(id, isCompleted)
}

func (u *TodoUsecase) DeleteTodo(id int) error {
	return u.repo.DeleteTodo(id)
}

package usecase

import "github.com/ShivangSrivastava/cli-todo/internal/entity"

type TodoRepository interface {
	GetTodos() ([]entity.Todo, error)
	CreateTodo(title string) (entity.Todo, error)
	UpdateTodoTitle(id int, title string) (entity.Todo, error)
	UpdateTodoIsCompleted(id int, isCompleted bool) (entity.Todo, error)
	DeleteTodo(id int) error
}

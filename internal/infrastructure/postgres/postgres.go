package postgres

import (
	"database/sql"
	"log"

	"github.com/ShivangSrivastava/cli-todo/internal/entity"
	_ "github.com/lib/pq"
)

type TodoRepo struct {
	DB *sql.DB
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todos(
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		is_completed BOOLEAN NOT NULL DEFAULT false
	)`
	_, err := db.Exec(query)
	return err
}

func NewDBRepo(db *sql.DB) *TodoRepo {
	err := createTable(db)
	if err != nil {
		log.Fatal(err)
	}
	return &TodoRepo{
		DB: db,
	}
}

// Create a todo
func (r *TodoRepo) CreateTodo(title string) error {
	query := `INSERT INTO todos (title) VALUES ($1)`
	_, err := r.DB.Exec(query, title)
	return err
}

// Get all todos
func (r *TodoRepo) GetTodos() ([]entity.Todo, error) {
	query := `SELECT id, title, is_completed FROM todos`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []entity.Todo
	for rows.Next() {
		var todo entity.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.IsCompleted)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// Update title of a todo
func (r *TodoRepo) UpdateTodoTitle(id int, title string) error {
	query := `UPDATE todos SET title = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, title, id)
	return err
}

// Update is_completed of a todo
func (r *TodoRepo) UpdateTodoIsCompleted(id int, isCompleted bool) error {
	query := `UPDATE todos SET is_completed = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, isCompleted, id)
	return err
}

// Delete a todo
func (r *TodoRepo) DeleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

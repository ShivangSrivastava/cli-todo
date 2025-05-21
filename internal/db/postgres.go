package db

import (
	"database/sql"
	"log"

	"github.com/ShivangSrivastava/cli-todo/internal/models"
	_ "github.com/lib/pq"
)

type DBRepo struct {
	DB *sql.DB
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todos(
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description VARCHAR(255),
		is_completed BOOLEAN NOT NULL DEFAULT false
	)`
	_, err := db.Exec(query)
	return err
}

func NewDBRepo(db *sql.DB) *DBRepo {
	err := createTable(db)
	if err != nil {
		log.Fatal(err)
	}
	return &DBRepo{
		DB: db,
	}
}

// Create a todo
func (r *DBRepo) CreateTodo(title, description string) error {
	query := `INSERT INTO todos (title, description) VALUES ($1, $2)`
	_, err := r.DB.Exec(query, title, description)
	return err
}

// Get all todos
func (r *DBRepo) GetTodos() ([]models.Todo, error) {
	query := `SELECT id, title, description, is_completed FROM todos`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// Update title of a todo
func (r *DBRepo) UpdateTodoTitle(id int, title string) error {
	query := `UPDATE todos SET title = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, title, id)
	return err
}

// Update description of a todo
func (r *DBRepo) UpdateTodoDescription(id int, description string) error {
	query := `UPDATE todos SET description = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, description, id)
	return err
}

// Update is_completed of a todo
func (r *DBRepo) UpdateTodoIsCompleted(id int, isCompleted bool) error {
	query := `UPDATE todos SET is_completed = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, isCompleted, id)
	return err
}

// Delete a todo
func (r *DBRepo) DeleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

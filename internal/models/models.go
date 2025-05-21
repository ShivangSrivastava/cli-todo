package models

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
}

type Todo struct {
	ID          int
	Title       string
	Description string
	IsCompleted bool
}

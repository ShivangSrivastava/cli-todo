package main

import (
	"database/sql"
	"log"

	"github.com/ShivangSrivastava/cli-todo/internal/config"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}
	dotenv := config.GetConfig()

	connStr := "postgres://" +
		dotenv.PostgresUser +
		":" + dotenv.PostgresPassword +
		"@localhost/" + dotenv.PostgresDB +
		"?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}

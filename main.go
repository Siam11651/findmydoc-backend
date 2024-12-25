package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms/ollama"

	_ "github.com/lib/pq"

	"findmydoc-backend/database"
	"findmydoc-backend/llm"
	"findmydoc-backend/routes"
)

func main() {
	{
		err := godotenv.Load()

		if err != nil {
			panic(err)
		}
	}

	r := gin.Default()

	{
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		db, err := sql.Open(
			"postgres",
			fmt.Sprintf(
				"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
				dbUser,
				dbPassword,
				dbName,
				dbHost,
				dbPort,
			),
		)

		if err != nil {
			panic(err)
		}

		database.Db = db
	}

	{
		ollama, err := ollama.New(ollama.WithModel("llama3.2:1b"))

		if err != nil {
			panic(err)
		}

		llm.Llm = ollama
	}

	routes.RegisterRoutes(r)
	r.Run(":3000")
}

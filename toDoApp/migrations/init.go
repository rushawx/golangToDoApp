package main

import (
	"fmt"
	"log"
	"os"
	"toDo/internal/todo"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%v user=%v password=%v dbname=%v port=%v",
				os.Getenv("PG_HOST"),
				os.Getenv("PG_USER"),
				os.Getenv("PG_PASSWORD"),
				os.Getenv("PG_DATABASE"),
				os.Getenv("PG_PORT"),
			),
		),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&todo.Task{})
	if err != nil {
		log.Fatal(err)
	}
}

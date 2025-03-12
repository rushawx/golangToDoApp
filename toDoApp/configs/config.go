package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Dsn string
}

func DefaultConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
	return &Config{
		Db: DbConfig{
			Dsn: fmt.Sprintf(
				"host=%v user=%v password=%v dbname=%v port=%v",
				os.Getenv("PG_HOST"),
				os.Getenv("PG_USER"),
				os.Getenv("PG_PASSWORD"),
				os.Getenv("PG_DATABASE"),
				os.Getenv("PG_PORT"),
			),
		},
	}
}

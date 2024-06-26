package initializer

import (
	"fmt"
	"log"
	"order/pkg/config"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func LoadDBConfig() *config.DatabaseConfig {

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("Incorrect Port Value, Setting to default 5432")
		port = 5432
	}

	return &config.DatabaseConfig{
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: port,
		DB_USER: os.Getenv("DB_USER"),
		DB_PSWD: os.Getenv("DB_PASSWORD"),
		DB_NAME: os.Getenv("DB_NAME"),
	}
}

package loadenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, string, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	typeOfStorage := os.Getenv("typeOfStorage")
	port := os.Getenv("port")
	pathToFile := os.Getenv("pathToFile")

	return typeOfStorage, port, pathToFile
}

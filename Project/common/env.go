package common

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Get port value from .env file
func GetPort() int {
	// Загрузка переменных окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Получение значения переменной окружения "PORT"
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT environment variable is not set")
	}

	// Преобразование строки в число
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port value: %v", err)
	}
	return port
}

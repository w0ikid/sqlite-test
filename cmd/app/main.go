package main

import (
	"github.com/joho/godotenv"
	"github.com/w0ikid/sqlite-test/internal"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env не найден или не загружен")
	}

	app.Run("config.yaml")
}

package main

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"github.com/w0ikid/sqlite-test/internal/configs"
	"github.com/w0ikid/sqlite-test/internal/connections"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env не найден или не загружен")
	}
}


func main() {
	// CONFIG {cleanenv or viper}

	// cleanenv
	loader := configs.CleanenvLoader{}

	// viper
	// loader := configs.ViperLoader{}

	cfg := configs.InitConfig(loader, "config.yaml")

	// CONNECTION
	connector, err := connections.GetConnector(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	db, err := connector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected successfully", db.Ping())



	fmt.Println(cfg)
}

package app

import (
	"fmt"
	"github.com/w0ikid/sqlite-test/internal/configs"
	"github.com/w0ikid/sqlite-test/internal/connections"
	"log"
)

func Run(path string) {
	// CONFIG {cleanenv or viper}

	// cleanenv
	loader := configs.CleanenvLoader{}

	// viper
	// loader := configs.ViperLoader{}

	cfg := configs.InitConfig(loader, path)

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

package main

import (
	"fmt"

	"github.com/w0ikid/sqlite-test/internal/configs"
)

func main() {
	loader := configs.CleanenvLoader{}
	// loader := configs.ViperLoader{}

	cfg := configs.InitConfig(loader, "config.yaml")

	fmt.Println(cfg)
}

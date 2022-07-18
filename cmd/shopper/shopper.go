package main

import (
	"flag"
	"log"

	"github.com/arturyumaev/shopper-api/config"
	"github.com/arturyumaev/shopper-api/internal/application"
	"github.com/arturyumaev/shopper-api/models"
)

func parseFlags() (*models.Flags, error) {
	flags := &models.Flags{}

	var configPath string
	flag.StringVar(&configPath, "config", "./config/config.yaml", "path to config file")
	flag.Parse()

	if err := config.ValidateConfigPath(configPath); err != nil {
		return nil, err
	}

	flags.ConfigPath = configPath

	return flags, nil
}

func main() {
	flags, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	config, err := config.ReadConfig(flags.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	app := application.NewApplication(config)

	if err = app.Run(); err != nil {
		log.Fatal(err)
	}
}

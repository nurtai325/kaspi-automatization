package config

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/nurtai325/kaspi/mailing/internal/models"
)

func Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	config = models.Config{}
	err = env.Parse(&config)
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	config.WORK_DIR = wd

	return nil
}

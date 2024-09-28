package config

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/nurtai325/kaspi/mailing/internal/models"
)

func New() (models.Config, error) {
	var config models.Config
	err := env.Parse(&config)
	if err != nil {
		return config, err
	}

	wd, err := os.Getwd()
	if err != nil {
		return config, err
	}
	config.WORK_DIR = wd

	return config, err
}

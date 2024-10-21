package config

import (
	"github.com/nurtai325/kaspi/mailing/internal/models"
)

var config models.Config

func New() models.Config {
	return config
}

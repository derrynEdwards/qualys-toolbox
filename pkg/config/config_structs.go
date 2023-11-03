package config

import (
	"github.com/derrynEdwards/qualys-toolbox/internal/qualysapi"
)

type Config struct {
	QualysApiClient qualysapi.Client
}

type ApiConfig struct {
	BaseURL  string
	User     string
	Password string
}

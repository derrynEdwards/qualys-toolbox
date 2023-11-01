package config

import (
	"github.com/derrynEdwards/qualys-toolbox/internal/qualysapi"
)

type Config struct {
	QualysApiClient qualysapi.Client
}

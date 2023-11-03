package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfigs() ApiConfig {
	godotenv.Load(".env")

	url := os.Getenv("QUALYS_API_URL")
	user := os.Getenv("QUALYS_API_USER")
	pass := os.Getenv("QUALYS_API_PASS")

	var apiCfg ApiConfig = ApiConfig{
		BaseURL:  url,
		User:     user,
		Password: pass,
	}

	return apiCfg
}

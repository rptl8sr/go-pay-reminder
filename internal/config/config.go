package config

import (
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

const (
	serviceAccountCreds = "serviceCreds.json"
)

type Config struct {
	Service  *jwt.Config
	TGToken  string     `env:"TG_API_TOKEN" env-required:"true"`
	ChatID   int        `env:"CHAT_ID" env-required:"true"`
	SheetID  string     `env:"SHEET_ID" env-required:"true"`
	LogLevel slog.Level `env:"LOG_LEVEL" env-required:"true"`
}

func MustLoad() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	if err = cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	f, err := os.ReadFile(serviceAccountCreds)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(f, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}

	cfg.Service = config

	return cfg, nil
}

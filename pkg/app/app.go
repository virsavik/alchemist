package app

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	AppConfigNameKey    = "APP_NAME"
	AppConfigVersionKey = "APP_VERSION"
	AppConfigPortKey    = "APP_PORT"
	AppConfigPGURLKey   = "APP_PG_URL"
)

type AppConfig struct {
	Name    string
	Version string
	Port    string
	PGURL   string
}

func ReadConfigFromEnv() AppConfig {
	return AppConfig{
		Name:    strings.TrimSpace(os.Getenv(AppConfigNameKey)),
		Version: strings.TrimSpace(os.Getenv(AppConfigVersionKey)),
		Port:    strings.TrimSpace(os.Getenv(AppConfigPortKey)),
		PGURL:   strings.TrimSpace(os.Getenv(AppConfigPGURLKey)),
	}
}

func (cfg AppConfig) IsValid() error {
	if cfg.Name == "" {
		return errors.New("name is required")
	}

	if cfg.Version == "" {
		return errors.New("version is required")
	}

	if cfg.Port == "" {
		return errors.New("port is required")
	}

	if num, err := strconv.Atoi(cfg.Port); err != nil || (num < 0 || num > 9999) {
		return errors.New("port is invalid")
	}

	if cfg.PGURL == "" {
		return errors.New("pgurl is required")
	}

	return nil
}

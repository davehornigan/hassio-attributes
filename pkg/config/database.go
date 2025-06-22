package config

import "fmt"

type DatabaseConfig struct {
	DatabaseHost string `env:"APP_DATABASE_HOST"`
	DatabasePort int    `env:"APP_DATABASE_PORT" envDefault:"5432"`
	DatabaseName string `env:"APP_DATABASE_NAME"`
	DatabaseUser string `env:"APP_DATABASE_USER"`
	DatabasePass string `env:"APP_DATABASE_PASS"`
}

func (cfg DatabaseConfig) GetDsnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DatabaseUser,
		cfg.DatabasePass,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
	)
}

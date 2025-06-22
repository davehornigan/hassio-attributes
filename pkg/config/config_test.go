package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"hassio-attributes/pkg/config"
)

type TestConfig struct {
	AppName string `env:"APP_NAME" envDefault:"MyApp"`
	Port    int    `env:"PORT" envDefault:"8080"`
	Secret  string `env:"SECRET" required:"true"`
}

func (c TestConfig) Validate() error {
	if len(c.Secret) < 8 {
		return &configError{"SECRET too short"}
	}
	return nil
}

type configError struct{ msg string }

func (e *configError) Error() string { return e.msg }

func TestLoadConfig(t *testing.T) {
	os.Clearenv()
	_ = os.Setenv("APP_NAME", "TestApp")
	_ = os.Setenv("PORT", "5050")
	_ = os.Setenv("SECRET", "supersecret")

	cfg := config.LoadConfig(TestConfig{})

	require.Equal(t, "TestApp", cfg.AppName)
	require.Equal(t, 5050, cfg.Port)
	require.Equal(t, "supersecret", cfg.Secret)
}

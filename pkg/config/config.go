package config

import (
	"fmt"
	"github.com/caarlos0/env"
	v "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
)

var dotenvFiles = []string{
	".env",
	".env.override",
	".env.local",
}

type Validatable interface {
	Validate() error
}

func LoadConfig[T interface{}](cfg T) T {
	for _, f := range dotenvFiles {
		if err := loadDotEnvFile(f); err != nil {
			log.Printf("⚠️ Skipped loading %s: %v", f, err)
		}
	}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("❌ Error parsing config \"%s\" from ENV: %s", reflect.TypeOf(cfg), err)
	}

	if validatable, ok := any(cfg).(Validatable); ok {
		if err := validatable.Validate(); err != nil {
			log.Fatalf("❌ Config \"%s\" validation failed: %s", reflect.TypeOf(cfg), err)
		}
	} else {
		validator := v.New(v.WithRequiredStructEnabled())
		if err := validator.Struct(cfg); err != nil {
			log.Fatalf("❌ Config \"%s\" validation failed: %s", reflect.TypeOf(cfg), err)
		}
	}

	log.Printf("✅ Config \"%s\" loaded\n", reflect.TypeOf(cfg))
	return cfg
}

func loadDotEnvFile(filename string) error {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip if file not exists
		}
		return fmt.Errorf("stat failed: %w", err)
	}
	if info.IsDir() {
		return fmt.Errorf("is a directory")
	}

	if err := godotenv.Overload(filename); err != nil {
		return fmt.Errorf("load failed: %w", err)
	}

	log.Printf("📄 Loaded %s", filename)
	return nil
}

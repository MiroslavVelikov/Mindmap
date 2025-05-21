package utils

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
)

const (
	postgres = "postgres"
	Logger   = "logger"

	UNIQUE_VIOLATION      = "23505"
	FOREIGN_KEY_VIOLATION = "23503"
	NOT_NULL_VIOLATION    = "23502"
	UNCLASIFIED_ERROR     = "unclassified error: %w"
)

type Config struct {
	DbUser     string `envconfig:"DB_USER"`
	DbPassword string `envconfig:"DB_PASSWORD"`
	DbName     string `envconfig:"DB_NAME"`
	DbHost     string `envconfig:"DB_HOST"`
	DbPort     string `envconfig:"DB_PORT"`
}

// This function has testing purposes
func onlyBackendRunning() Config {
	return Config{
		DbUser:     "postgres",
		DbPassword: "postgres",
		DbName:     "mindmapTEST",
		DbHost:     "localhost",
		DbPort:     "5432",
	}
}

func GetConnectionString() (string, error) {
	var cfg Config

	// because of testing purpouse
	if true {
		cfg = onlyBackendRunning()
	} else {
		if err := envconfig.Process("", &cfg); err != nil {
			return "", err
		}
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	fmt.Println(connectionString)

	return connectionString, nil
}

func ConnectToDB() (*sqlx.DB, error) {
	connectionString, err := GetConnectionString()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect(postgres, connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

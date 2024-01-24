package config

import (
	"fmt"
	"os"
)

func MustGet(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("env variable %s is not set", key)
	}
	return value, nil
}

type Configuration struct {
	Endpoint string
	Database string
	Username string
	Password string
}

func GetConfiguration() (*Configuration, error) {
	endpoint, err := MustGet("ENDPOINT")
	if err != nil {
		return nil, err
	}

	database, err := MustGet("DATABASE")
	if err != nil {
		return nil, err
	}

	username, err := MustGet("USERNAME")
	if err != nil {
		return nil, err
	}
	password, err := MustGet("PASSWORD")
	if err != nil {
		return nil, err
	}

	return &Configuration{
		Endpoint: endpoint,
		Database: database,
		Username: username,
		Password: password,
	}, nil
}

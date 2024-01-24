package pkg

import (
	"context"
	"fmt"
	"lindorm-cli/pkg/client"
	"log"
	"os"
)

func MustGet(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("env variable %s is not set", value)
	}
	return value, nil
}

func Invoke(statement string) {
	resp, err := query(statement)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	render(resp)
}

func render(resp *client.QueryResponse) error {
	fmt.Println(*resp)
	return nil
}

func query(statement string) (*client.QueryResponse, error) {
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
	password, err := MustGet("password")
	if err != nil {
		return nil, nil, err
	}
	client := client.NewClient(
		endpoint,
		database,
		username,
		password,
	)

	context := context.Background()

	return client.Query(context, statement)
}

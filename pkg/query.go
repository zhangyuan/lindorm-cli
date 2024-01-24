package query

import (
	"context"
	"fmt"
	"lindorm-cli/pkg/client"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
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

func Invoke(statement string) error {
	conf, err := getConf()
	if err != nil {
		return err
	}

	resp, err := query(conf, statement)
	if err != nil {
		return err
	}

	render(resp)

	return nil
}

func getConf() (*Configuration, error) {
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

func render(resp *client.QueryResponse) error {
	writer := table.NewWriter()
	header := table.Row{}
	for _, column := range resp.Columns {
		header = append(header, column)
	}
	writer.AppendHeader(header)

	for _, row := range resp.Rows {
		for idx, field := range row {
			if resp.Metadata[idx] == "TIMESTAMP" {
				if fieldAsInt64, ok := field.(float64); ok {
					timestampAsTime := time.UnixMilli(int64(fieldAsInt64))
					row[idx] = timestampAsTime.Local().Format(time.RFC3339)
				}
			}
		}
		writer.AppendRow(row)
	}

	fmt.Println(writer.Render())

	return nil
}

func query(conf *Configuration, statement string) (*client.QueryResponse, error) {
	client := client.NewClient(
		conf.Endpoint,
		conf.Database,
		conf.Username,
		conf.Password,
	)

	context := context.Background()

	return client.Query(context, statement)
}

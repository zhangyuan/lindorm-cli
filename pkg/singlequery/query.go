package singlequery

import (
	"context"
	"lindorm-cli/pkg/client"
	"lindorm-cli/pkg/config"
	"lindorm-cli/pkg/render"
)

func Invoke(statement string) error {
	conf, err := config.GetConfiguration()
	if err != nil {
		return err
	}

	resp, err := query(conf, statement)
	if err != nil {
		return err
	}

	render.Render(resp)

	return nil
}

func query(conf *config.Configuration, statement string) (*client.QueryResponse, error) {
	client := client.NewClient(
		conf.Endpoint,
		conf.Database,
		conf.Username,
		conf.Password,
	)

	context := context.Background()

	return client.Query(context, statement)
}

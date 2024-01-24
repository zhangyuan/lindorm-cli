package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	httpClient *resty.Client
	endpoint   string
	database   string
	username   string
	password   string
}

func NewClient(endpoint string, database string, username string, password string) *Client {
	client := resty.New()

	return &Client{
		httpClient: client,
		endpoint:   endpoint,
		database:   database,
		username:   username,
		password:   password,
	}
}

type QueryResponse struct {
	Columns  []string        `json:"columns"`
	Metadata []string        `json:"metadata"`
	Rows     [][]interface{} `json:"rows"`
}

type QueryErrorResponse struct {
	Code     int    `json:"code"`
	SqlState string `json:"sqlstate"`
	Message  string `json:"message"`
}

func (client *Client) Query(context context.Context, statement string) (*QueryResponse, error) {
	uri := fmt.Sprintf("%s/api/v2/sql", strings.TrimRight(client.endpoint, "/"))

	var queryResponse QueryResponse

	resp, err := client.httpClient.R().SetContext(context).EnableTrace().
		SetBasicAuth(client.username, client.password).
		SetQueryParam("database", client.database).
		SetQueryParam("chunked", "false").
		SetBody(statement).
		SetResult(&queryResponse).
		Post(uri)

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.String())
	}

	if err != nil {
		return nil, err
	}

	return &queryResponse, nil
}

package render

import (
	"fmt"
	"lindorm-cli/pkg/client"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func Render(resp *client.QueryResponse) error {
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

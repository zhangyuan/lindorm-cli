package cmd

import (
	"lindorm-cli/pkg/singlequery"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var statement string
var statementFile string

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query the data via SQL statements",
	Run: func(cmd *cobra.Command, args []string) {
		if statementFile != "" {
			bytes, err := os.ReadFile(statementFile)
			if err != nil {
				log.Fatalf("%v", err)
			}
			if err := singlequery.Invoke(string(bytes)); err != nil {
				log.Fatalf("%v", err)
			}
		} else {
			if err := singlequery.Invoke(statement); err != nil {
				log.Fatalf("%v", err)
			}
		}
	},
}

func init() {
	queryCmd.Flags().StringVarP(&statement, "statement", "s", "", "SQL Statement to query")
	queryCmd.Flags().StringVarP(&statementFile, "file", "f", "", "Fie path containing SQL statement to query")
	queryCmd.MarkFlagsMutuallyExclusive("statement", "file")
	queryCmd.MarkFlagsOneRequired("statement", "file")

	rootCmd.AddCommand(queryCmd)
}

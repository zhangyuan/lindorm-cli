package cmd

import (
	"lindorm-cli/pkg/prompt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lindorm-cli",
	Short: "A lindorm-cli with less bugs",
	Run: func(cmd *cobra.Command, args []string) {
		if err := prompt.Prompt(); err != nil {
			log.Fatalf("error: %v", err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

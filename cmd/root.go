package cmd

import (
	"lindorm-cli/pkg/interactivequery"
	"lindorm-cli/pkg/version"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/subosito/gotenv"
)

var rootCmd = &cobra.Command{
	Use:     "lindorm-cli",
	Short:   "A lindorm-cli with less bugs",
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {
		if err := gotenv.Load(); err != nil {
			log.Fatalf("error: %v", err)
		}

		// https: //github.com/c-bata/go-prompt/issues/228
		defer handleExit()

		if err := interactivequery.RunPrompt(); err != nil {
			log.Fatalf("error: %v\n", err)
		}
	},
}

func handleExit() {
	rawModeOff := exec.Command("/bin/stty", "-raw", "echo")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

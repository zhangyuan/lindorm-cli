package interactivequery

import (
	"context"
	"fmt"
	"lindorm-cli/pkg/client"
	"lindorm-cli/pkg/config"
	"lindorm-cli/pkg/render"
	"strings"

	"github.com/c-bata/go-prompt"
)

const exitCommand = "exit"

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func RunPrompt() error {
	configuration, err := config.GetConfiguration()
	if err != nil {
		return err
	}

	client := client.NewClient(
		configuration.Endpoint,
		configuration.Database,
		configuration.Username,
		configuration.Password,
	)

	context := context.Background()

	isExitCommand := func(in string) bool {
		return strings.TrimRightFunc(in, func(r rune) bool {
			return r == ';' || r == ' '
		}) == exitCommand
	}

	executor := func(in string) {
		if isExitCommand(in) {
			return
		}

		resp, err := client.Query(context, in)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			if err := render.Render(resp); err != nil {
				fmt.Printf("error: %v\n", err)
			}
		}
	}

	exitCheckOnInput := func(in string, breakline bool) bool {
		return isExitCommand(in) && breakline
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(configuration.Database+" > "),
		prompt.OptionSetExitCheckerOnInput(exitCheckOnInput),
	)
	p.Run()

	fmt.Println("Bye!")

	return nil
}

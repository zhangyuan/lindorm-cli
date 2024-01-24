package interactivequery

import (
	"context"
	"lindorm-cli/pkg/client"
	"lindorm-cli/pkg/config"
	"lindorm-cli/pkg/render"
	"log"
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

	executor := func(in string) {
		resp, err := client.Query(context, in)
		if err != nil {
			log.Printf("error: %v", err)
		} else {
			if err := render.Render(resp); err != nil {
				log.Printf("error: %v", err)
			}
		}
	}

	exitCheckOnInput := func(in string, breakline bool) bool {
		return strings.TrimRightFunc(in, func(r rune) bool {
			return r == ';' || r == ' '
		}) == exitCommand && breakline
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(configuration.Database+" > "),
		prompt.OptionSetExitCheckerOnInput(exitCheckOnInput),
	)
	p.Run()

	return nil
}

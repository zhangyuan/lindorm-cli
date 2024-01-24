package prompt

import (
	"context"
	"lindorm-cli/pkg/client"
	"lindorm-cli/pkg/config"
	"lindorm-cli/pkg/render"
	"log"

	"github.com/manifoldco/promptui"
)

func Prompt() error {
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

	prompt := promptui.Prompt{
		Label:       "SQL",
		HideEntered: false,
	}

	for {
		input, err := prompt.Run()
		if err != nil {
			return err
		}

		context := context.Background()

		resp, err := client.Query(context, input)

		if err != nil {
			log.Printf("error: %v\n", err)
			continue
		}

		render.Render(resp)
	}
}

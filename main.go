package main

import (
	"lindorm-cli/cmd"
	"log"

	"github.com/subosito/gotenv"
)

func main() {
	if err := gotenv.Load(); err != nil {
		log.Fatalf("error: %v", err)
	}
	cmd.Execute()
}

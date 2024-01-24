package main

import (
	"lindorm-cli/cmd"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	cmd.Execute()
}

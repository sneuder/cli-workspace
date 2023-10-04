package main

import (
	"workspace/internal/cmd"
	"workspace/internal/config"
)

func main() {
	config.Config()
	cmd.StartTerminal()
}

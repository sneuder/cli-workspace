package cmd

import (
	"bufio"
	"os"
	"strings"
	"workspace/internal/cmd/basics"
	"workspace/internal/cmd/workspace"
)

type ActionCMD func([]string)

var actionsCMD = map[string]map[string]ActionCMD{
	"workspace": {
		"workspace": workspace.DecribeCMD,
		"create":    workspace.Create,
		"run":       workspace.Run,
		"build":     workspace.Build,
		"stop":      workspace.Stop,
		"rm":        workspace.Remove,
		"ls":        workspace.Ls,
	},
	"clear": {
		"clear": basics.Clear,
	},
	"version": {
		"version": basics.Version,
	},
	"help": {
		"help": basics.Help,
	},
	"exit": {
		"exit": basics.Exit,
	},
}

func StartTerminal() {

	reader := bufio.NewReader(os.Stdin)

	for {
		print("workspace > ")

		input := receiveInput(reader)
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		inputs := fromStringToArray(input)

		receiveAction(inputs)
	}
}

func receiveInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')

	if err != nil {
		println("error reading input:", err)
		return "error reading input"
	}

	return input
}

func receiveAction(actionKeys []string) {

	if len(actionKeys) == 1 {
		actionKeys = append(actionKeys, actionKeys[0])
	}

	storeActions := actionsCMD[actionKeys[0]]
	action, exists := storeActions[actionKeys[1]]

	if !exists {
		println("command not found")
		return
	}

	action(actionKeys[2:])
}

func fromStringToArray(str string) []string {
	words := strings.Split(str, " ")
	return words
}

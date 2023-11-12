package cmd

import (
	"bufio"
	"os"
	"strings"
	"workspace/internal/cmd/basics"
	"workspace/internal/cmd/dockerFast/dfBuild"
	"workspace/internal/cmd/dockerFast/dfRun"
	"workspace/internal/cmd/dockerFast/dfStart"
	"workspace/internal/cmd/workspace/wsBuild"
	"workspace/internal/cmd/workspace/wsLs"
	"workspace/internal/cmd/workspace/wsRemove"
	"workspace/internal/cmd/workspace/wsRun"
	"workspace/internal/cmd/workspace/wsStop"
	"workspace/internal/cmd/workspace/wsUtil"
	"workspace/internal/constants"
)

type ActionCMD func([]string)

var actionsCMD = map[string]map[string]ActionCMD{
	string(constants.ActionWorkspace): {
		"workspace": wsUtil.DecribeCMD,
		"run":       wsRun.Run,
		"build":     wsBuild.Build,
		"stop":      wsStop.Stop,
		"rm":        wsRemove.Remove,
		"ls":        wsLs.Ls,
	},
	string(constants.ActionDocker): {
		"run":   dfRun.Run,
		"start": dfStart.Start,
		"build": dfBuild.Build,
	},
	string(constants.ActionClear): {
		"clear": basics.Clear,
	},
	string(constants.ActionVersion): {
		"version": basics.Version,
	},
	string(constants.ActionHelp): {
		"help": basics.Help,
	},
	string(constants.ActionExit): {
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

		inputs := strings.Split(input, " ")
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

	if !validateArgs(actionKeys, actionKeys[2:]) {
		return
	}

	action(actionKeys[2:])
}

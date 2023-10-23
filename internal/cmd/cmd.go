package cmd

import (
	"bufio"
	"os"
	"strings"
	"workspace/internal/cmd/basics"
	"workspace/internal/cmd/workspace/wsBuild"
	"workspace/internal/cmd/workspace/wsCreate"
	"workspace/internal/cmd/workspace/wsLs"
	"workspace/internal/cmd/workspace/wsRemove"
	"workspace/internal/cmd/workspace/wsRun"
	"workspace/internal/cmd/workspace/wsStop"
	"workspace/internal/cmd/workspace/wsUtil"
	"workspace/internal/constants"
	"workspace/internal/util"
)

type ActionCMD func([]string)

var subActionsToValidateArgs = []string{
	string(constants.SubActionWSCreate),
	string(constants.SubActionWSStop),
	string(constants.SubActionWSRun),
	string(constants.SubActionWSBuild),
	string(constants.SubActionWSRemove),
}

var actionsCMD = map[string]map[string]ActionCMD{
	string(constants.ActionWorkspace): {
		"workspace": wsUtil.DecribeCMD,
		"create":    wsCreate.Create,
		"run":       wsRun.Run,
		"build":     wsBuild.Build,
		"stop":      wsStop.Stop,
		"rm":        wsRemove.Remove,
		"ls":        wsLs.Ls,
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

	existArgToValidate := util.ContainsString(subActionsToValidateArgs, actionKeys[1])
	existsArgAction := len(actionKeys) >= 3

	if existArgToValidate && !existsArgAction {
		message := "workspace name required"

		if string(constants.SubActionWSBuild) == actionKeys[1] {
			message = "workspace.json path required"
		}

		println(message)
		return
	}

	action(actionKeys[2:])
}

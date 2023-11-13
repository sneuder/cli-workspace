package cmd

import (
	"fmt"
	"workspace/internal/constants"
)

type ValidationMessages struct {
	WorkspaceName     string
	ImageName         string
	ContainerName     string
	BinMountPath      string
	PathWorkspaceJson string
}

var validationMesssages = ValidationMessages{
	WorkspaceName:     "workspace name required",
	ImageName:         "image name required",
	ContainerName:     "container name required",
	BinMountPath:      "bind path required",
	PathWorkspaceJson: "workspace.json path required",
}

var validationArgs = map[string]map[string]map[int]string{
	string(constants.ActionDocker): {
		"run": {
			0: validationMesssages.ImageName,
			1: validationMesssages.ContainerName,
			2: validationMesssages.BinMountPath,
		},
		"start": {
			0: validationMesssages.ContainerName,
		},
		"stop": {
			0: validationMesssages.ContainerName,
		},
		"remove": {
			0: validationMesssages.ContainerName,
		},
	},
	string(constants.ActionWorkspace): {
		"run": {
			0: validationMesssages.WorkspaceName,
		},
		"build": {
			0: validationMesssages.PathWorkspaceJson,
		},
		"rm": {
			0: validationMesssages.WorkspaceName,
		},
		"stop": {
			0: validationMesssages.WorkspaceName,
		},
	},
}

func validateArgs(actionKeys []string, args []string) bool {
	validationArg := validationArgs[actionKeys[0]][actionKeys[1]]
	lenArgs := len(args)
	messageError, exists := validationArg[lenArgs]

	if exists {
		fmt.Println(messageError)
		return false
	}

	return true
}

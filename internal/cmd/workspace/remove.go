package workspace

import (
	"fmt"
	"path"
	"workspace/internal/config"
	"workspace/internal/docker/container"
	"workspace/internal/docker/image"
	"workspace/internal/file"
)

type ActionSequence struct {
	Action func(string)
	State  State
}

var actionsSequence = []ActionSequence{
	{
		Action: container.Stop,
		State:  Running,
	},
	{
		Action: container.Remove,
		State:  Built,
	},
	{
		Action: image.Remove,
		State:  Instanced,
	},
	{
		Action: removeFile,
		State:  Inactive,
	},
}

func Remove(args []string) {

	if len(args) == 0 {
		fmt.Println("workspace name needed")
		return
	}

	workspaceName := args[0]
	containerState := getState(workspaceName)

	if containerState == Nonexistent {
		fmt.Println("workspace does not exist")
		return
	}

	sequeceConnected := false
	for _, actionSequence := range actionsSequence {
		if actionSequence.State == containerState {
			sequeceConnected = true
		}

		if !sequeceConnected {
			continue
		}

		actionSequence.Action(workspaceName)
	}
}

func removeFile(fileName string) {
	filePathWorkspace := path.Join(config.PathDirs["workspaces"], fileName+"-workspace")
	filePathConfig := path.Join(config.PathDirs["workspaces"], fileName+"-config")
	file.Remove(filePathWorkspace)
	file.Remove(filePathConfig)
}

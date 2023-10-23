package wsRemove

import (
	"fmt"
	"path"
	"workspace/internal/cmd/workspace/wsData"
	"workspace/internal/cmd/workspace/wsUtil"
	"workspace/internal/config"
	"workspace/internal/docker/container"
	"workspace/internal/docker/image"
	"workspace/internal/file"
)

type ActionSequence struct {
	Action func(string)
	State  wsData.State
}

var actionsSequence = []ActionSequence{
	{
		Action: container.Stop,
		State:  wsData.Running,
	},
	{
		Action: container.Remove,
		State:  wsData.Built,
	},
	{
		Action: image.Remove,
		State:  wsData.Instanced,
	},
	{
		Action: removeFile,
		State:  wsData.Inactive,
	},
}

func Remove(args []string) {
	workspaceName := args[0]
	containerState := wsUtil.GetState(workspaceName)

	if containerState == wsData.Nonexistent {
		fmt.Println("workspace does not exist")
		return
	}

	fmt.Println("removing workspace...")

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

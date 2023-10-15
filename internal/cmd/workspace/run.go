package workspace

import (
	"fmt"
	"path"
	"workspace/internal/config"
	"workspace/internal/docker/container"
	"workspace/internal/docker/image"
	"workspace/internal/file"
	"workspace/internal/util"
)

func Run(args []string) {

	if len(args) == 0 {
		fmt.Println("workspace name needed")
		return
	}

	workspaceName := args[0]
	containerState := getState(workspaceName)

	if containerState == Nonexistent {
		fmt.Println("workspace does not exists")
		return
	}

	// if the container exists and we do not have to built it again
	if containerState == Built {
		container.Run(workspaceName)
		return
	}

	// if someone remove the image
	if containerState == Inactive {
		rebuildImage(workspaceName)
	}

	// build the container process

	dataContainer["name"] = workspaceName

	contentFile := file.Read(path.Join(config.PathDirs["workspaces"], dataContainer["name"]+"-config"))
	contentFileMap := util.StringToMap(contentFile, "=")

	dataContainer["bindmount"] = contentFileMap["BINDMOUNTPATH"]
	dataContainer["ports"] = contentFileMap["EXPOSEPORTS"]

	container.StartContainerProcess(dataContainer)
}

func rebuildImage(workspaceName string) {
	file.Rename(workspaceName+"-workspace", "dockerfile")
	image.BuildImage(workspaceName)
	file.Rename("dockerfile", workspaceName+"-workspace")
}

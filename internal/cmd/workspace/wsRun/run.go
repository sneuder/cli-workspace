package wsRun

import (
	"fmt"
	"path"
	"workspace/internal/cmd/workspace/wsData"
	"workspace/internal/cmd/workspace/wsUtil"
	"workspace/internal/config"
	"workspace/internal/docker/container"
	"workspace/internal/docker/image"
	"workspace/internal/file"
	"workspace/internal/util"
)

func Run(args []string) {
	workspaceName := args[0]
	containerState := wsUtil.GetState(workspaceName)

	if containerState == wsData.Running {
		fmt.Println("workspace already running")
		return
	}

	if containerState == wsData.Nonexistent {
		fmt.Println("workspace does not exists")
		return
	}

	// if the container exists and we do not have to built it again
	if containerState == wsData.Built {
		container.Run(workspaceName)
		return
	}

	// if someone remove the image
	if containerState == wsData.Inactive {
		rebuildImage(workspaceName)
	}

	// build the container process
	fmt.Println("starting workspace...")
	wsData.DataContainer["name"] = workspaceName

	fullPathWSConfig := path.Join(config.PathDirs["workspaces"], wsData.DataContainer["name"]+"-config")
	contentFile := file.Read(fullPathWSConfig)
	contentFileMap := util.StringToMap(contentFile, "=")

	wsData.DataContainer["bindmount"] = contentFileMap["BINDMOUNTPATH"]
	wsData.DataContainer["ports"] = contentFileMap["EXPOSEPORTS"]
	wsData.DataContainer["networks"] = contentFileMap["NETWORKS"]

	container.Create(wsData.DataContainer)
	resetDataContainer()
	fmt.Println("workspace running")
}

func rebuildImage(workspaceName string) {
	file.Rename(workspaceName+"-workspace", "dockerfile")
	image.Build(workspaceName)
	file.Rename("dockerfile", workspaceName+"-workspace")
}

func resetDataContainer() {
	for key := range wsData.DataContainer {
		wsData.DataContainer[key] = ""
	}
}

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

	// can continues based on state
	if !canContinue(containerState) {
		return
	}

	if !canContinueBasedOnAction(workspaceName, containerState) {
		return
	}

	// build the container process
	fmt.Println("starting workspace")

	wsData.DataContainer["name"] = workspaceName
	getDataFromFileConfig()
	container.Create(wsData.DataContainer)

	resetDataContainer()
}

func rebuildImage(workspaceName string) {
	file.Rename(workspaceName+"-workspace", "dockerfile")
	image.Build(workspaceName)
	file.Rename("dockerfile", workspaceName+"-workspace")
}

func getDataFromFileConfig() {
	fullPathWSConfig := path.Join(config.PathDirs["workspaces"], wsData.DataContainer["name"]+"-config")
	contentFile, _ := file.Read(fullPathWSConfig)
	contentFileMap := util.StringToMap(contentFile, "=")

	wsData.DataContainer["bindmount"] = contentFileMap["BINDMOUNTPATH"]
	wsData.DataContainer["ports"] = contentFileMap["EXPOSEPORTS"]
	wsData.DataContainer["networks"] = contentFileMap["NETWORKS"]
}

func canContinue(containerState wsData.State) bool {
	if containerState == wsData.Running {
		fmt.Println("workspace already running")
		return false
	}

	if containerState == wsData.Nonexistent {
		fmt.Println("workspace does not exists")
		return false
	}

	return true
}

func canContinueBasedOnAction(workspaceName string, containerState wsData.State) bool {
	// if the container exists and we do not have to built it again
	if containerState == wsData.Built {
		fmt.Println("starting workspace")
		container.Run(workspaceName)
		return false
	}

	// if someone remove the image
	if containerState == wsData.Inactive {
		fmt.Println("building workspace")
		rebuildImage(workspaceName)
		return true
	}

	return true
}

func resetDataContainer() {
	for key := range wsData.DataContainer {
		wsData.DataContainer[key] = ""
	}
}

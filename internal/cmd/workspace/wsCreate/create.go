package wsCreate

import (
	"fmt"
	"path"
	"workspace/internal/cmd/workspace/wsData"
	"workspace/internal/config"
	"workspace/internal/docker/image"
	"workspace/internal/file"
)

func Create(args []string) {
	if len(args) != 0 && validateExistance(args[0]) {
		fmt.Println("workspace name already exists")
		return
	}

	fmt.Println("creating workspace")

	// seting data
	setArgs(args)

	// build process
	image.Create(wsData.DataWorkspace)
	setConfigFile()

	// reseting date
	resetWorkspaceData()
}

func validateExistance(workspaceName string) bool {
	filePathWorkspace := path.Join(config.PathDirs["workspaces"], workspaceName+"-workspace")
	exists := image.Exists(workspaceName) || file.FileExists(filePathWorkspace)
	return exists
}

func setConfigFile() {
	fileName := wsData.DataWorkspace["name"].Value + "-config"
	file.Open(fileName, config.PathDirs["workspaces"])

	file.Write([]byte("BINDMOUNTPATH=" + wsData.DataWorkspace["bindmount"].Value))
	file.Write([]byte("EXPOSEPORTS=" + wsData.DataWorkspace["ports"].Value))
	file.Write([]byte("NETWORKS=" + wsData.DataWorkspace["networks"].Value))

	file.Close()
}

func setArgs(args []string) {
	if len(args) == 0 {
		return
	}

	workspaceName := args[0]

	data := wsData.DataWorkspace["name"]
	data.Value = workspaceName
	wsData.DataWorkspace["name"] = data
}

func resetWorkspaceData() {
	for key, itemData := range wsData.DataWorkspace {
		itemData.Value = ""
		wsData.DataWorkspace[key] = itemData
	}
}

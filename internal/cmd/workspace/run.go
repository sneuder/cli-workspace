package workspace

import (
	"fmt"
	"path"
	"workspace/internal/config"
	"workspace/internal/docker/container"
	"workspace/internal/file"
	"workspace/internal/util"
)

func Run(args []string) {

	if len(args) == 0 {
		fmt.Println("workspace name needed")
		return
	}

	containerState := getState(args[0])

	if containerState == Nonexistent {
		fmt.Println("workspace does not exists")
		return
	}

	// if the container exists and we do not have to built it again
	if containerState == Built {
		container.Run(args[0])
		return
	}

	dataContainer["name"] = args[0]

	contentFile := file.Read(path.Join(config.PathDirs["workspaces"], dataContainer["name"]+"-config"))
	contentFileMap := util.StringToMap(contentFile, "=")

	dataContainer["bindMount"] = contentFileMap["BINDMOUNTPATH"]
	dataContainer["ports"] = contentFileMap["EXPOSEPORTS"]

	container.StartContainerProcess(dataContainer)
}

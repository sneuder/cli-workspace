package wsStop

import (
	"fmt"
	"workspace/internal/cmd/workspace/wsData"
	"workspace/internal/cmd/workspace/wsUtil"
	"workspace/internal/docker/container"
)

func Stop(args []string) {
	workspaceName := args[0]
	containerState := wsUtil.GetState(workspaceName)

	if containerState == wsData.Nonexistent {
		fmt.Println("workspace does not exists")
		return
	}

	if containerState != wsData.Running {
		fmt.Println("workspace is not in running state")
		return
	}

	fmt.Println("stopping workspace...")
	container.Stop(workspaceName)
}

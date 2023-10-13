package workspace

import (
	"fmt"
	"workspace/internal/docker/container"
)

func Stop(args []string) {
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

	if containerState != Running {
		fmt.Println("workspace is not in running state")
		return
	}

	container.Stop(workspaceName)
}

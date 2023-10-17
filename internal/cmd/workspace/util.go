package workspace

import (
	"fmt"
	"workspace/internal/docker"
)

type State string

const (
	Inactive    State = "inactive"
	Instanced   State = "instanced"
	Built       State = "built"
	Running     State = "running"
	Nonexistent State = "nonexistent"
)

func getState(workspaceName string) State {
	containerInfo, _ := docker.GetContainerInfo(workspaceName)

	exists := validateExistance(workspaceName)

	if !exists {
		return Nonexistent
	}

	if containerInfo.ID == "" {
		return Inactive
	}

	if containerInfo.State.Status == "" {
		return Instanced
	}

	if containerInfo.State.Status == "exited" {
		return Built
	}

	if containerInfo.State.Status == "running" {
		return Running
	}

	return Instanced
}

func DecribeCMD(args []string) {
	fmt.Println("usage: workspace")
	fmt.Printf("  %-20s- %s\n", "create", "Create a workspace.")
	fmt.Printf("  %-20s- %s\n", "build", "Build from workspace.json")
	fmt.Printf("  %-20s- %s\n", "run", "Run a workspace")
	fmt.Printf("  %-20s- %s\n", "stop", "Stop a workspace")
	fmt.Printf("  %-20s- %s\n", "ls", "List all workspaces")
	fmt.Printf("  %-20s- %s\n", "remove", "Remove a workspace.")
}

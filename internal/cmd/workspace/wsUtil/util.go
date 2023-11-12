package wsUtil

import (
	"fmt"
	"path"
	"workspace/internal/cmd/workspace/wsData"
	"workspace/internal/config"
	"workspace/internal/docker"
	"workspace/internal/docker/image"
	"workspace/internal/file"
)

func GetState(workspaceName string) wsData.State {
	containerInfo, _ := docker.GetContainerInfo(workspaceName)
	exists := WorkspaceExists(workspaceName)

	if !exists {
		return wsData.Nonexistent
	}

	if containerInfo.ID == "" {
		return wsData.Inactive
	}

	if containerInfo.State.Status == "" {
		return wsData.Instanced
	}

	if containerInfo.State.Status == "exited" {
		return wsData.Built
	}

	if containerInfo.State.Status == "running" {
		return wsData.Running
	}

	return wsData.Instanced
}

func WorkspaceExists(workspaceName string) bool {
	filePathWorkspace := path.Join(config.PathDirs["workspaces"], workspaceName+"-workspace")
	exists := image.Exists(workspaceName) || file.FileExists(filePathWorkspace)
	return exists
}

func DecribeCMD(args []string) {
	fmt.Println("usage: workspace")
	fmt.Printf("  %-20s- %s\n", "build", "Build from workspace.json")
	fmt.Printf("  %-20s- %s\n", "run", "Run a workspace")
	fmt.Printf("  %-20s- %s\n", "stop", "Stop a workspace")
	fmt.Printf("  %-20s- %s\n", "ls", "List all workspaces")
	fmt.Printf("  %-20s- %s\n", "remove", "Remove a workspace.")
}

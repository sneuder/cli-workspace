package docker

import (
	"encoding/json"
	"os/exec"
	"workspace/internal/cmd/workspace/wsData"
	"workspace/internal/model"
)

func GetContainerInfo(containerName string) (model.ContainerInfo, error) {
	cmd := exec.Command("docker", "inspect", containerName)
	output, err := cmd.Output()

	if err != nil {
		return model.ContainerInfo{}, err
	}

	var containerInfo []model.ContainerInfo
	err = json.Unmarshal(output, &containerInfo)

	if err != nil {
		return model.ContainerInfo{}, err
	}

	return containerInfo[0], nil
}

func ExistsElements(elementType string, elementName string) bool {
	cmd := exec.Command("docker", elementType, "inspect", elementName)
	_, err := cmd.Output()
	return err == nil
}

func GetState(containerName string) wsData.State {
	containerInfo, err := GetContainerInfo(containerName)

	if err != nil {
		return wsData.Nonexistent
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

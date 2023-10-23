package docker

import (
	"encoding/json"
	"os/exec"
	"workspace/internal/model"
)

func GetContainerInfo(workspaceName string) (model.ContainerInfo, error) {
	cmd := exec.Command("docker", "inspect", workspaceName)
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

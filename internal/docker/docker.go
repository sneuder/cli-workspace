package docker

import (
	"encoding/json"
	"os/exec"
)

type State struct {
	Status     string `json:"Status"`
	Running    bool   `json:"Running"`
	Paused     bool   `json:"Paused"`
	Restarting bool   `json:"Restarting"`
	OOMKilled  bool   `json:"OOMKilled"`
	Dead       bool   `json:"Dead"`
	Pid        int    `json:"Pid"`
	ExitCode   int    `json:"ExitCode"`
	Error      string `json:"Error"`
	StartedAt  string `json:"StartedAt"`
	FinishedAt string `json:"FinishedAt"`
}

type ContainerInfo struct {
	ID      string `json:"Id"`
	Created string `json:"Created"`
	Path    string `json:"Path"`
	State   State  `json:"State"`
}

func GetContainerInfo(workspaceName string) (ContainerInfo, error) {
	cmd := exec.Command("docker", "inspect", workspaceName)
	output, err := cmd.Output()

	if err != nil {
		return ContainerInfo{}, err
	}

	var containerInfo []ContainerInfo
	err = json.Unmarshal(output, &containerInfo)

	if err != nil {
		return ContainerInfo{}, err
	}

	return containerInfo[0], nil
}

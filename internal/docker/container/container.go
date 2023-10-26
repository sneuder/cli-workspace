package container

import (
	"log"
	"os/exec"
	"path"
	"strings"
	"workspace/internal/config"
)

var buildContainerCMD = []string{}

func Create(dataContainer map[string]string) {
	setInitializer()

	setBindMount(dataContainer["bindmount"])
	setExposePort(dataContainer["ports"])
	setNetworks(dataContainer["networks"])
	setContainerName(dataContainer["name"])

	buildContainer()
	resetContainerCMD()
}

func Run(containerName string) {
	cmd := exec.Command("docker", "start", containerName)
	cmd.Output()
}

func Stop(workspaceName string) {
	cmd := exec.Command("docker", "container", "stop", workspaceName)
	cmd.Output()
}

func Exists(containerName string) bool {
	cmd := exec.Command("docker", "container", "inspect", containerName)
	_, err := cmd.Output()
	return err == nil
}

func Remove(workspaceName string) {
	cmd := exec.Command("docker", "container", "rm", workspaceName)
	cmd.Output()
}

//...

func buildContainer() {
	cmd := exec.Command(buildContainerCMD[0], buildContainerCMD[1:]...)
	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}

func setInitializer() {
	buildContainerCMD = append(buildContainerCMD, "docker", "run", "-d")
}

func setContainerName(workspaceName string) {
	buildContainerCMD = append(buildContainerCMD, "--name", workspaceName, workspaceName)
}

func setExposePort(exposePorts string) {
	if exposePorts == "" {
		return
	}

	collectionPort := strings.Split(exposePorts, " ")

	for _, port := range collectionPort {
		buildContainerCMD = append(buildContainerCMD, "-p", port+":"+port)
	}
}

func setBindMount(pathBindMount string) {
	fullPathBindMount := path.Join(config.BasePath, pathBindMount)
	bindMountPartCMD := `type=bind,source=` + fullPathBindMount + `,target=/workspace`
	buildContainerCMD = append(buildContainerCMD, "--mount", bindMountPartCMD)
}

func setNetworks(networks string) {
	if networks == "" {
		return
	}

	collectionNetwork := strings.Split(networks, " ")

	for _, network := range collectionNetwork {
		buildContainerCMD = append(buildContainerCMD, "--network="+network)
	}
}

func resetContainerCMD() {
	buildContainerCMD = []string{}
}

package container

import (
	"log"
	"os/exec"
	"path"
	"strings"
	"workspace/internal/config"
)

var buildContainerCMD = []string{}

func StartContainerProcess(dataContainer map[string]string) {
	setInitializer()

	setBindMount(dataContainer["bindmount"])
	setExposePort(dataContainer["ports"])
	setContainerName(dataContainer["name"])

	buildContainer()
}

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

func Run(workspaceName string) {
	cmd := exec.Command("docker", "start", workspaceName)
	cmd.Output()
}

func Stop(workspaceName string) {
	cmd := exec.Command("docker", "container", "stop", workspaceName)
	cmd.Output()
}

func Remove(workspaceName string) {
	cmd := exec.Command("docker", "container", "rm", workspaceName)
	cmd.Output()
}

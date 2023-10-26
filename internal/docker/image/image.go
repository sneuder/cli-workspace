package image

import (
	"log"
	"os/exec"
	"strings"
	"workspace/internal/config"
	"workspace/internal/file"
	"workspace/internal/model"
)

func Create(dataWorkspace map[string]model.ValuesWorkspace) {
	file.Open("dockerfile", config.PathDirs["workspaces"])

	setImage(dataWorkspace["image"].Value)
	setUpdate()
	setTools(dataWorkspace["tools"].Value)
	setWorkDir()
	setPorts(dataWorkspace["ports"].Value)

	setCMD()

	file.Close()
	Build(dataWorkspace["name"].Value)
	file.Rename("dockerfile", dataWorkspace["name"].Value+"-workspace")
}

func Build(imageName string) {
	cmd := exec.Command("docker", "build", "-t", imageName, config.PathDirs["workspaces"])
	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}

func Remove(workspaceName string) {
	cmd := exec.Command("docker", "image", "rm", workspaceName)
	cmd.Output()
}

func Exists(imageName string) bool {
	cmd := exec.Command("docker", "image", "inspect", imageName)
	_, err := cmd.Output()
	return err == nil
}

//...

func setImage(value string) {
	image := "FROM " + value
	file.Write([]byte(image))
}

func setUpdate() {
	updata := "RUN apt-get update"
	file.Write([]byte(updata))
}

func setTools(tools string) {
	if tools == "" {
		return
	}

	toolToInstall := "RUN apt install " + tools + " -y"
	file.Write([]byte(toolToInstall))
}

func setWorkDir() {
	workdir := "WORKDIR /workspace"
	file.Write([]byte(workdir))
}

func setPorts(ports string) {
	if ports == "" {
		return
	}

	collectionPort := strings.Split(ports, " ")

	for _, port := range collectionPort {
		portToSet := "EXPOSE " + port
		file.Write([]byte(portToSet))
	}
}

func setCMD() {
	cmd := `CMD ["sleep", "infinity"]`
	file.Write([]byte(cmd))
}

package docker

import (
	"log"
	"os/exec"
	"workspace/internal/config"
	"workspace/internal/file"
	"workspace/internal/model"
)

func StartImageProcess(dataWorkspace map[string]model.DataWorkspace) {
	file.Open("dockerfile", config.PathDirs["workspaces"])

	setImage(dataWorkspace["image"].Value)
	setUpdate()
	setTools(dataWorkspace["tools"].Value)
	setWorkDir()

	setCMD()

	file.Close()
	buildImage(dataWorkspace["name"].Value)
	file.Rename("dockerfile", dataWorkspace["name"].Value)
}

func setImage(value string) {
	image := "FROM " + value
	file.Write([]byte(image))
}

func setUpdate() {
	updata := "RUN apt-get update"
	file.Write([]byte(updata))
}

func setTools(tools string) {
	toolToInstall := "RUN apt install " + tools + " -y"
	file.Write([]byte(toolToInstall))
}

func setWorkDir() {
	workdir := "WORKDIR /workspace"
	file.Write([]byte(workdir))
}

func setCMD() {
	cmd := `CMD ["sleep", "infinity"]`
	file.Write([]byte(cmd))
}

func buildImage(imageName string) {
	cmd := exec.Command("docker", "build", "-t", imageName, config.PathDirs["workspaces"])

	output, err := cmd.Output()

	outputStr := string(output)

	if err != nil {
		log.Fatal(err)
	}

	println(outputStr)
}

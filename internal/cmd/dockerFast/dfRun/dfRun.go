package dfRun

import (
	"fmt"
	"os/exec"
	"path"
	"workspace/internal/config"
)

func Run(args []string) {
	imageName := args[0]
	containerName := args[1]
	bindMountPath := args[2]

	fmt.Println("building image")
	buildContainer(imageName, containerName, bindMountPath)
}

func buildContainer(imageName string, containerName string, bindMountPath string) {
	bindMountCmd := setBindMount(bindMountPath)
	cmd := exec.Command("docker", "run", "-d", "--mount", bindMountCmd, "--name", containerName, imageName)
	cmd.Output()
}

func setBindMount(pathBindMount string) string {
	fullPathBindMount := path.Join(config.BasePath, pathBindMount)
	bindMountPartCMD := `type=bind,source=` + fullPathBindMount + `,target=/workspace`
	return bindMountPartCMD
}

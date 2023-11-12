package dfStop

import (
	"fmt"
	"workspace/internal/docker/container"
)

func Stop(args []string) {
	containerName := args[0]
	fmt.Println("stopping container")
	container.Stop(containerName)
}

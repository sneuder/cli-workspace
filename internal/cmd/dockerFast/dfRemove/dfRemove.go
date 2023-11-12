package dfRemove

import (
	"fmt"
	"workspace/internal/docker/container"
)

func Remove(args []string) {
	containerName := args[0]
	fmt.Println("removing container")
	container.Stop(containerName)
}

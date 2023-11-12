package dfStart

import (
	"fmt"
	"workspace/internal/docker/container"
)

func Start(args []string) {
	containerName := args[0]
	fmt.Println("starting container")
	container.Run(containerName)
}

package dfBuild

import (
	"fmt"
	"workspace/internal/docker/image"
)

func Build(args []string) {
	imageName := args[0]
	imagePath := args[1]

	fmt.Println("building image")
	image.Build(imageName, imagePath)
}

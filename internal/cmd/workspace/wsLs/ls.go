package wsLs

import (
	"fmt"
	"os"
	"strings"
	"workspace/internal/cmd/workspace/wsUtil"
	"workspace/internal/config"
)

func Ls(args []string) {
	folderToRead := config.PathDirs["workspaces"]
	files, _ := os.ReadDir(folderToRead)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		if !strings.HasSuffix(fileName, "workspace") {
			continue
		}

		fileName = strings.Replace(fileName, "-workspace", "", -1)
		fmt.Printf("- %-10s %s\n", fileName, wsUtil.GetState(fileName))
	}
}

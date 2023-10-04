package workspace

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"workspace/internal/config"
	"workspace/internal/docker"
	"workspace/internal/file"
	"workspace/internal/model"
	"workspace/internal/util"
)

var orderToGetData = []string{"name", "image", "tools", "bindMount"}

var dataWorkspace = map[string]model.DataWorkspace{
	"name": {
		Text:     "Name for workspace: ",
		Value:    "",
		Required: true,
	},
	"image": {
		Text:     "Image for workspace: ",
		Value:    "",
		Required: true,
	},
	"tools": {
		Text:     "Tools to include: ",
		List:     []string{"git", "make"},
		Value:    "",
		Required: false,
	},
	"bindMount": {
		Text:     "Path for workspace: ",
		Value:    "",
		Required: true,
	},
}

var dataContainer = map[string]string{
	"name":      "",
	"bindMount": "",
	"port":      "",
}

func Create(args []string) {
	getDataWorkspace()
	docker.StartImageProcess(dataWorkspace)
	setConfigFile()
}

func Run(args []string) {
	dataContainer["name"] = args[0]

	contentFile := file.Read(path.Join(config.PathDirs["workspaces"], dataContainer["name"]+"-config"))
	contentFileMap := util.StringToMap(contentFile)

	dataContainer["bindMount"] = contentFileMap["BINDMOUNTPATH"]

	docker.StartContainerProcess(dataContainer)
}

func Ls(args []string) {

}

func Stop(args []string) {

}

func Remove(args []string) {

}

func DecribeCMD(a []string) {
	fmt.Println("usage: workspace")
	fmt.Printf("  %-20s- %s\n", "create", "Create a workspace.")
	fmt.Printf("  %-20s- %s\n", "run", "Run a workspace")
	fmt.Printf("  %-20s- %s\n", "stop", "Stop a workspace")
	fmt.Printf("  %-20s- %s\n", "ls", "List all workspaces")
	fmt.Printf("  %-20s- %s\n", "remove", "Remove a workspace.")
}

func setConfigFile() {
	fileName := dataWorkspace["name"].Value + "-config"
	file.Open(fileName, config.PathDirs["workspaces"])
	file.Write([]byte("BINDMOUNTPATH=" + dataWorkspace["bindMount"].Value))
	file.Close()
}

func getDataWorkspace() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(orderToGetData); i++ {
		data := dataWorkspace[orderToGetData[i]]
		fmt.Print(data.Text)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if data.Required && input == "" {
			i -= 1
			continue
		}

		data.Value = input
		dataWorkspace[orderToGetData[i]] = data
	}
}

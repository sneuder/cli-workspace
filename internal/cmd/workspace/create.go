package workspace

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"workspace/internal/config"
	"workspace/internal/docker/image"
	"workspace/internal/docker/network"
	"workspace/internal/file"
)

func Create(args []string) {
	if len(args) == 0 {
		fmt.Println("workspace name needed")
		return
	}

	if len(args) != 0 && validateExistance(args[0]) {
		fmt.Println("workspace name already exists")
		return
	}

	fmt.Println("creating workspace...")

	// seting data
	setArgs(args)
	getDataWorkspace()

	// build process
	// createNetwork()
	image.StartImageProcess(dataWorkspace)
	setConfigFile()

	// reseting date
	resetWorkspaceData()
}

func validateExistance(workspaceName string) bool {
	filePathWorkspace := path.Join(config.PathDirs["workspaces"], workspaceName+"-workspace")
	exists := image.Exists(workspaceName) || file.FileExists(filePathWorkspace)
	return exists
}

func getDataWorkspace() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(orderToGetData); i++ {
		data := dataWorkspace[orderToGetData[i]]

		if data.FullFilled {
			continue
		}

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

func setConfigFile() {
	fileName := dataWorkspace["name"].Value + "-config"
	file.Open(fileName, config.PathDirs["workspaces"])

	file.Write([]byte("BINDMOUNTPATH=" + dataWorkspace["bindmount"].Value))
	file.Write([]byte("EXPOSEPORTS=" + dataWorkspace["ports"].Value))
	file.Write([]byte("NETWORKS=" + dataWorkspace["networks"].Value))

	file.Close()
}

func setArgs(args []string) {
	if len(args) == 0 {
		return
	}

	workspaceName := args[0]

	data := dataWorkspace["name"]
	data.Value = workspaceName
	dataWorkspace["name"] = data
}

func createNetwork() {
	if dataWorkspace["networks"].Value == "" {
		return
	}

	collectionNetwork := strings.Split(dataWorkspace["networks"].Value, " ")

	for _, networkName := range collectionNetwork {
		if network.Exists(networkName) {
			continue
		}

		network.Create(networkName)
	}
}

func resetWorkspaceData() {
	for key, itemData := range dataWorkspace {
		itemData.Value = ""
		dataWorkspace[key] = itemData
	}
}

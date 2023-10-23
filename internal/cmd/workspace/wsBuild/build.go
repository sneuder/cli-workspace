package wsBuild

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"workspace/internal/cmd/workspace/wsCreate"
	"workspace/internal/cmd/workspace/wsData"
	"workspace/internal/cmd/workspace/wsUtil"
	"workspace/internal/config"
	"workspace/internal/file"
	"workspace/internal/model"
	"workspace/internal/util"
)

var configWorkspace model.ConfigWorkspace

func Build(args []string) {
	if len(args) == 0 {
		fmt.Println("workspace path needed")
		return
	}

	configWorkspace = getWorkspaceConfig(args)
	workspaceState := wsUtil.GetState(configWorkspace.Name)

	if workspaceState != wsData.Nonexistent {
		fmt.Println("workspace already exist, change name in workspace.json")
		return
	}

	fmt.Println("building workspace...")

	configWorkspace.BindMount = util.JoinPathArgs(args)
	setWokspaceInfo()

	args[0] = configWorkspace.Name

	wsCreate.Create(args)
	resetConfigWorkspace()
}

func getWorkspaceConfig(args []string) model.ConfigWorkspace {
	workspacePath := getPathWorkSpaceInfo(args)
	dataContent := file.Read(workspacePath)

	var configWorkspace model.ConfigWorkspace
	json.Unmarshal([]byte(dataContent), &configWorkspace)

	return configWorkspace
}

func setWokspaceInfo() {
	wsData.DataWorkspace["name"] = model.ValuesWorkspace{
		Value:      configWorkspace.Name,
		FullFilled: true,
	}

	wsData.DataWorkspace["image"] = model.ValuesWorkspace{
		Value:      configWorkspace.Image,
		FullFilled: true,
	}

	wsData.DataWorkspace["bindmount"] = model.ValuesWorkspace{
		Value:      configWorkspace.Image,
		FullFilled: true,
	}

	toolsValue := strings.Join(configWorkspace.Tools, " ")
	wsData.DataWorkspace["tools"] = model.ValuesWorkspace{
		Value:      toolsValue,
		FullFilled: true,
	}

	networksValue := strings.Join(configWorkspace.Networks, " ")
	wsData.DataWorkspace["networks"] = model.ValuesWorkspace{
		Value:      networksValue,
		FullFilled: true,
	}
}

// func setDatabaseInfo() {
// 	for _, configDatabase := range configWorkspace.Databases {
// 		configInformation := map[string]model.DataConfig{
// 			"name": {
// 				Text:       "Name: ",
// 				Value:      configDatabase.Name,
// 				Required:   true,
// 				FullFilled: false,
// 			},
// 			"type": {
// 				Text:       "Database Type: ",
// 				Value:      configDatabase.Type,
// 				Required:   true,
// 				FullFilled: false,
// 			},
// 			"networks": {
// 				Text:       "Networks: ",
// 				Value:      strings.Join(configDatabase.Networks, " "),
// 				Required:   false,
// 				FullFilled: false,
// 			},
// 			"ports": {
// 				Text:       "Exposed ports: ",
// 				Value:      strings.Join(configDatabase.Ports, " "),
// 				Required:   false,
// 				FullFilled: false,
// 			},
// 		}

// 		dataDBs = append(dataDBs, configInformation)
// 	}
// }

// func createDatabase() {
// 	setDatabaseInfo()
// 	for _, dataDB := range dataDBs {
// 		database.ExistsOrCreate(dataDB)
// 	}
// }

func getPathWorkSpaceInfo(args []string) string {
	argPath := util.JoinPathArgs(args)
	return path.Join(config.BasePath, argPath, "workspace.json")
}

func resetConfigWorkspace() {
	// configWorkspace = un
}

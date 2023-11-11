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
	"workspace/internal/docker/database"
	"workspace/internal/file"
	"workspace/internal/model"
	"workspace/internal/util"
)

var configWorkspace model.ConfigWorkspace

func Build(args []string) {
	workspaceConfigError := getWorkspaceConfig(args)

	if !canContinue(workspaceConfigError) {
		return
	}

	fmt.Println("building workspace")
	configWorkspace.BindMount = util.JoinPathArgs(args)
	setWokspaceInfo()

	args[0] = configWorkspace.Name

	wsCreate.Create(args)
	resetConfigWorkspace()
}

func getWorkspaceConfig(args []string) error {
	workspacePath := getPathWorkSpaceInfo(args)
	dataContent, err := file.Read(workspacePath)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(dataContent), &configWorkspace)

	return nil
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
		Value:      configWorkspace.BindMount,
		FullFilled: true,
	}

	toolsValue := strings.Join(configWorkspace.Tools, " ")
	wsData.DataWorkspace["tools"] = model.ValuesWorkspace{
		Value:      toolsValue,
		FullFilled: true,
	}

	portsValue := strings.Join(configWorkspace.Ports, " ")
	wsData.DataWorkspace["ports"] = model.ValuesWorkspace{
		Value:      portsValue,
		FullFilled: true,
	}

	networksValue := strings.Join(configWorkspace.Networks, " ")
	wsData.DataWorkspace["networks"] = model.ValuesWorkspace{
		Value:      networksValue,
		FullFilled: true,
	}
}

func canContinue(workspaceConfigError error) bool {
	if workspaceConfigError != nil {
		fmt.Println("workspace.json does not exist")
		return false
	}

	if util.ContainsUpperCases(configWorkspace.Name) {
		fmt.Println("workspace name has to be in lowercase")
		return false
	}

	workspaceState := wsUtil.GetState(configWorkspace.Name)
	if workspaceState != wsData.Nonexistent {
		fmt.Println("workspace already exists")
		return false
	}

	return true
}

func buildDatabase() {
	fmt.Println("building databases...")
	if len(configWorkspace.Databases) == 0 {
		return
	}

	for _, db := range configWorkspace.Databases {
		database.Build(db)
	}
}

func getPathWorkSpaceInfo(args []string) string {
	argPath := util.JoinPathArgs(args)
	return path.Join(config.BasePath, argPath, "workspace.json")
}

func resetConfigWorkspace() {
	configWorkspace = model.ConfigWorkspace{}
}

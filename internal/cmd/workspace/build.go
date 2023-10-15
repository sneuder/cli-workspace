package workspace

import (
	"encoding/json"
	"fmt"
	"path"
	"reflect"
	"strings"
	"workspace/internal/config"
	"workspace/internal/file"
	"workspace/internal/util"
)

type ConfigWorkspace struct {
	Name      string   `json:"name"`
	Image     string   `json:"image"`
	BindMount string   `json:"bindMount"`
	Tools     []string `json:"tools"`
	Ports     []string `json:"ports"`
}

func Build(args []string) {
	if len(args) == 0 {
		fmt.Println("Workspace path needed")
		return
	}

	configWorkspace := getWorkspaceInfo(args)
	workspaceState := getState(configWorkspace.Name)

	if workspaceState != Nonexistent {
		fmt.Println("Workspace already exist. Change name in workspace.json")
		return
	}

	configWorkspace.BindMount = util.JoinPathArgs(args)
	setWokspaceInfo(configWorkspace)

	args[0] = configWorkspace.Name

	fmt.Println("Building workspace...")

	Create(args)
}

func getWorkspaceInfo(args []string) ConfigWorkspace {
	workspacePath := getPathWorkSpaceInfo(args)
	dataContent := file.Read(workspacePath)

	var configWorkspace ConfigWorkspace
	json.Unmarshal([]byte(dataContent), &configWorkspace)

	return configWorkspace
}

func setWokspaceInfo(configWorkspace ConfigWorkspace) {
	val := reflect.ValueOf(configWorkspace)
	typ := reflect.TypeOf(configWorkspace)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fieldValue := field.Interface()

		data := dataWorkspace[strings.ToLower(fieldName)]
		value := fieldValue

		if reflect.TypeOf(value).Kind() == reflect.Slice {
			if len(value.([]string)) == 0 {
				value = ""
			} else {
				value = strings.Join(value.([]string)[:], " ")
			}
		}

		data.Value = value.(string)
		data.FullFilled = true
		dataWorkspace[strings.ToLower(fieldName)] = data
	}
}

func getPathWorkSpaceInfo(args []string) string {
	argPath := util.JoinPathArgs(args)
	return path.Join(config.BasePath, argPath, "workspace.json")
}

package database

import (
	"fmt"
	"os/exec"
	"workspace/internal/cmd/workspace/wsData"
	"workspace/internal/docker"
	"workspace/internal/docker/container"
	"workspace/internal/model"
)

var buildDatabaseCMD = []string{}

func Build(configWorkspaceDB model.ConfigWorkspaceDB) {
	databaseState := docker.GetState(configWorkspaceDB.Name)

	if !typeExists(configWorkspaceDB.Type) {
		fmt.Println("db " + configWorkspaceDB.Type + " does not exist")
		return
	}

	if databaseState == wsData.Running || databaseState == wsData.Built {
		fmt.Println("db " + configWorkspaceDB.Name + " already exists")
		return
	}

	buildDB(configWorkspaceDB)
	buildDatabaseCMD = []string{}
}

func Run(databaseName string) {
	fmt.Println("starting " + databaseName + " db")
	container.Run(databaseName)
}

func Stop(databaseName string) {
	fmt.Println("stopping " + databaseName + " db")
	container.Stop(databaseName)
}

func Remove(databaseName string) {
	fmt.Println("removing " + databaseName + " db")
	container.Remove(databaseName)
}

func buildDB(configWorkspaceDB model.ConfigWorkspaceDB) {
	fmt.Println("creating " + configWorkspaceDB.Name + " db can take a while")

	setInitializer()
	setDatabaseName(configWorkspaceDB.Name)
	setNetworks(configWorkspaceDB.Networks)
	setEnv(envBasedType[configWorkspaceDB.Type])
	setExposedPorts(configWorkspaceDB.Ports)
	setDatabaseType(configWorkspaceDB.Type)

	cmd := exec.Command(buildDatabaseCMD[0], buildDatabaseCMD[1:]...)

	_, err := cmd.Output()

	if err != nil {
		fmt.Println("something went wrong")
	}
}

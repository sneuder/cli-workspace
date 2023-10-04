package config

import (
	"os"
	"path"
	"runtime"
)

var BasePath string

var PathDirs = map[string]string{
	"workspace":  ".workspace",
	"workspaces": ".workspace/workspaces",
	"config":     ".workspace/config",
}

var Dirs = map[string]string{
	"workspace":  ".workspace",
	"workspaces": "workspaces",
	"config":     "config",
}

func Config() {
	setVariablesOS()
	buildDirs()
}

func setVariablesOS() {
	BasePath = "~"

	if runtime.GOOS == "windows" {
		BasePath = os.Getenv("UserProfile")
	}
}

func buildDirs() {
	for mapKey, pathDir := range PathDirs {
		fullPath := path.Join(BasePath, pathDir)
		buildFullPathDirs(mapKey, fullPath)
		os.Mkdir(fullPath, 0755)
	}
}

func buildFullPathDirs(dirKey string, pathValue string) {
	PathDirs[dirKey] = pathValue
}

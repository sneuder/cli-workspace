package workspace

import (
	"workspace/internal/model"
)

var orderToGetData = []string{"image", "tools", "ports", "bindMount"}

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
	"ports": {
		Text:     "Exposed ports: ",
		Value:    "",
		Required: false,
	},
}

var dataContainer = map[string]string{
	"name":      "",
	"bindMount": "",
	"ports":     "",
}

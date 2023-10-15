package workspace

import (
	"workspace/internal/model"
)

var orderToGetData = []string{"image", "tools", "ports", "bindmount"}

var dataWorkspace = map[string]model.DataWorkspace{
	"name": {
		Text:       "Name for workspace: ",
		Value:      "",
		Required:   true,
		FullFilled: false,
	},
	"image": {
		Text:       "Image for workspace: ",
		Value:      "",
		Required:   true,
		FullFilled: false,
	},
	"tools": {
		Text:       "Tools to include: ",
		Value:      "",
		Required:   false,
		FullFilled: false,
	},
	"bindmount": {
		Text:       "Path for workspace: ",
		Value:      "",
		Required:   true,
		FullFilled: false,
	},
	"ports": {
		Text:       "Exposed ports: ",
		Value:      "",
		Required:   false,
		FullFilled: false,
	},
}

var dataContainer = map[string]string{
	"name":      "",
	"bindmount": "",
	"ports":     "",
}

package wsData

import (
	"workspace/internal/model"
)

var OrderToGetData = []string{"image", "tools", "ports", "networks", "bindmount"}

var DataWorkspace = map[string]model.ValuesWorkspace{
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
	"networks": {
		Text:       "Networks: ",
		Value:      "",
		Required:   false,
		FullFilled: false,
	},
}

var DataContainer = map[string]string{
	"name":      "",
	"bindmount": "",
	"ports":     "",
}

var DataDB = map[string]model.ValuesWorkspace{
	"name": {
		Text:       "Name: ",
		Value:      "",
		Required:   true,
		FullFilled: false,
	},
	"type": {
		Text:       "Database Type: ",
		Value:      "",
		Required:   true,
		FullFilled: false,
	},
	"networks": {
		Text:       "Networks: ",
		Value:      "",
		Required:   false,
		FullFilled: false,
	},
	"ports": {
		Text:       "Exposed ports: ",
		Value:      "",
		Required:   false,
		FullFilled: false,
	},
}

var DataDBs = []map[string]model.ValuesWorkspace{}

type State string

const (
	Inactive    State = "inactive"
	Instanced   State = "instanced"
	Built       State = "built"
	Running     State = "running"
	Nonexistent State = "nonexistent"
)

package model

type ValuesWorkspace struct {
	Text       string
	Value      string
	Required   bool
	FullFilled bool
}

type ConfigWorkspace struct {
	Name      string   `json:"name"`
	Image     string   `json:"image"`
	BindMount string   `json:"bindMount"`
	Networks  []string `json:"networks"`
	Tools     []string `json:"tools"`
	Ports     []string `json:"ports"`
}

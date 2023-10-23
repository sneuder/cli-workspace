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

type ContainerState struct {
	Status     string `json:"Status"`
	Running    bool   `json:"Running"`
	Paused     bool   `json:"Paused"`
	Restarting bool   `json:"Restarting"`
	OOMKilled  bool   `json:"OOMKilled"`
	Dead       bool   `json:"Dead"`
	Pid        int    `json:"Pid"`
	ExitCode   int    `json:"ExitCode"`
	Error      string `json:"Error"`
	StartedAt  string `json:"StartedAt"`
	FinishedAt string `json:"FinishedAt"`
}

type ContainerInfo struct {
	ID      string         `json:"Id"`
	Created string         `json:"Created"`
	Path    string         `json:"Path"`
	State   ContainerState `json:"State"`
}

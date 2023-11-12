package constants

type Action string

const (
	ActionWorkspace Action = "workspace"
	ActionDocker    Action = "docker"
	ActionClear     Action = "clear"
	ActionVersion   Action = "version"
	ActionHelp      Action = "help"
	ActionExit      Action = "exit"
)

const (
	SubActionWSCreate Action = "create"
	SubActionWSRun    Action = "run"
	SubActionWSBuild  Action = "build"
	SubActionWSStop   Action = "stop"
	SubActionWSRemove Action = "rm"
	SubActionWSList   Action = "ls"
)

const (
	SubActionDfBuild Action = "build"
	SubActionDfRun   Action = "run"
	SubActionDfStop  Action = "stop"
	SubActionDfRm    Action = "rm"
	SubActionDfLs    Action = "ls"
)

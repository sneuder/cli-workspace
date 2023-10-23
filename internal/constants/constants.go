package constants

type Action string

const (
	ActionWorkspace Action = "workspace"
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

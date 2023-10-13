package workspace

import (
	"fmt"
)

func DecribeCMD(args []string) {
	fmt.Println("usage: workspace")
	fmt.Printf("  %-20s- %s\n", "create", "Create a workspace.")
	fmt.Printf("  %-20s- %s\n", "run", "Run a workspace")
	fmt.Printf("  %-20s- %s\n", "stop", "Stop a workspace")
	fmt.Printf("  %-20s- %s\n", "ls", "List all workspaces")
	fmt.Printf("  %-20s- %s\n", "remove", "Remove a workspace.")
}

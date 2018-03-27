package command

import "fmt"

type Help struct {
}

func (h *Help) Init(args []string) bool {
	return true
}

func (h *Help) Execute() {
	fmt.Println()
	fmt.Println("Usage:	gossh COMMAND")
	fmt.Println()
	fmt.Println("Options：")
	fmt.Println("  -v, -Version     ", "\t", "Print Version information and quit")
	fmt.Println()
	fmt.Println("Commands：")
	fmt.Println("  conn         ", "\t", "Connect to the server")
	fmt.Println("  ls           ", "\t", "Show the server list")
	fmt.Println("  add          ", "\t", "Add a server")
	fmt.Println("  rm           ", "\t", "Remove a server")
	fmt.Println()
}

func (h *Help) Help() {
}

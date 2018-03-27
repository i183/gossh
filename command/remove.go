package command

import (
	"fmt"

	"github.com/i183/gossh/server"
)

type Remove struct {
	serverName string //Server name
}

func (rm *Remove) Init(args []string) bool {
	if len(args) == 1 {
		rm.serverName = args[0]
		return true
	}
	return false
}

func (rm *Remove) Execute() {
	if ok := server.RemoveByName(rm.serverName); !ok {
		panic(fmt.Sprintf("Did not find the \"%s\" server", rm.serverName))
	}
}

func (rm *Remove) Help() {
	fmt.Println()
	fmt.Println("\"gossh rm\" requires at least 1 argument.")
	fmt.Println()
	fmt.Println("Usage:	gossh rm <SERVER_NAME>")
	fmt.Println()
}

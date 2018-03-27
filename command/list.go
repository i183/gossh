package command

import (
	"fmt"

	"github.com/i183/gossh/server"
)

type List struct {
}

func (ls *List) Init(args []string) bool {
	return true
}

func (ls *List) Execute() {
	servers := server.ReadAll()
	for _, sv := range servers {
		fmt.Println(sv.ServerName)
	}
}

func (ls *List) Help() {
}

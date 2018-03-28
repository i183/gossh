package main

import (
	"fmt"
	"os"

	"github.com/i183/gossh/command"
	"github.com/i183/gossh/server"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("gossh failed:", err)
		}
	}()

	server.InitServerFile() //Init Server file

	var cmd string
	var args []string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
		args = os.Args[2:]
	}
	h := command.CreateHandler(cmd) //create handler by command
	if ok := h.Init(args); ok {
		h.Execute()
	} else {
		h.Help()
	}

}

package main

import (
	"fmt"
	"os"
)

const (
	serverFileName = ".gossh"
	versionNumber  = "0.1"
	githubUrl      = "https://github.com/i183/gossh"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("gossh failed:", err)
		}
	}()

	initServerFile()                //Init Server file
	h := createHandler(os.Args[1:]) //create handler by args
	h.execute()                     //call execute function
}

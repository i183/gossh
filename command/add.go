package command

import (
	"fmt"

	"strconv"

	"github.com/i183/gossh/kit"
	"github.com/i183/gossh/server"
)

type Add struct {
	serverName string //Server name
	username   string //Username
	ip         string //IP Address
	port       int    //Port
	password   string //Password
}

func (add *Add) Init(args []string) bool {
	if len(args) == 5 {
		port, err := strconv.Atoi(args[3])
		kit.Check(err)

		add.serverName = args[0]
		add.username = args[1]
		add.ip = args[2]
		add.port = port
		add.password = args[4]
		return true
	}
	return false
}

func (add *Add) Execute() {
	sv := server.Server{
		ServerName: add.serverName,
		Username:   add.username,
		IP:         add.ip,
		Port:       add.port,
		Password:   add.password,
	}

	server.Add(sv)
}

func (add *Add) Help() {
	fmt.Println()
	fmt.Println("\"gossh add\" requires at least 5 argument.")
	fmt.Println()
	fmt.Println("Usage:	gossh add <SERVER_NAME> <USERNAME> <IP> <PORT> <PASSWORD>")
	fmt.Println()
}

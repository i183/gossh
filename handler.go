package main

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type server struct {
	serverName string //Server name
	username   string //Username
	ip         string //IP Address
	port       string //Port
	password   string //Password
}

type handler interface {
	execute()
}

type connection struct {
	serverName string //Server name
}

type add struct {
	serverName string //Server name
	username   string //Username
	ip         string //IP Address
	port       string //Port
	password   string //Password
}

type remove struct {
	serverName string //Server name
}

type list struct {
}

type version struct {
}

type help struct {
}

func (conn *connection) execute() {
	sv, ok := findServerByServerName(conn.serverName)
	if !ok {
		panic("Can not find server")
	}
	config := &ssh.ClientConfig{
		User: sv.username,
		Auth: []ssh.AuthMethod{
			ssh.Password(sv.password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", sv.ip+":"+sv.port, config)
	check(err)
	defer client.Close()

	session, err := client.NewSession()
	check(err)
	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	check(err)

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)
	check(err)

	defer terminal.Restore(fd, oldState)

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty("xterm-256color", termHeight, termWidth, modes)
	check(err)

	err = session.Shell()
	check(err)

	err = session.Wait()
	check(err)
}

func (add *add) execute() {
	sv := server{
		serverName: add.serverName,
		username:   add.username,
		ip:         add.ip,
		port:       add.port,
		password:   add.password,
	}

	addServer(sv)
}

func (rm *remove) execute() {
	removeServerByServerName(rm.serverName)
}

func (ls *list) execute() {
	servers := readAllServer()
	for _, sv := range servers {
		fmt.Println(sv.serverName)
	}
}

func (v *version) execute() {
	fmt.Println("Version: ", VERSION)
	fmt.Println("Github:", GITHUB_URL)
}

func (h *help) execute() {
	// TODO implement me
}

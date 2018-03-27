package command

import (
	"fmt"
	"net"
	"os"

	"strconv"

	"github.com/i183/gossh/kit"
	"github.com/i183/gossh/server"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type Connect struct {
	serverName string //Server name
}

func (conn *Connect) Init(args []string) bool {
	if len(args) == 1 {
		conn.serverName = args[0]
		return true
	}
	return false
}

func (conn *Connect) Execute() {
	sv, ok := server.FindByName(conn.serverName)
	if !ok {
		panic(fmt.Sprintf("Did not find the \"%s\" server", conn.serverName))
	}
	config := &ssh.ClientConfig{
		User: sv.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(sv.Password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", sv.IP+":"+strconv.Itoa(sv.Port), config)
	kit.Check(err)
	defer client.Close()

	session, err := client.NewSession()
	kit.Check(err)
	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	kit.Check(err)

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)
	kit.Check(err)

	defer terminal.Restore(fd, oldState)

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty("xterm-256color", termHeight, termWidth, modes)
	kit.Check(err)

	err = session.Shell()
	kit.Check(err)

	err = session.Wait()
	kit.Check(err)
}

func (conn *Connect) Help() {
	fmt.Println()
	fmt.Println("\"gossh conn\" requires at least 1 argument.")
	fmt.Println()
	fmt.Println("Usage:	gossh conn <SERVER_NAME>")
	fmt.Println()
}

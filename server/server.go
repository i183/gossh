package server

import (
	"bytes"
	"encoding/json"
	"os"
	"os/user"

	"github.com/i183/gossh/kit"
)

const serverFileName = ".gossh"

type Server struct {
	ServerName string //Server name
	Username   string //Username
	IP         string //IP Address
	Port       int    //Port
	Password   string //Password
}

func InitServerFile() {
	if _, err := os.Stat(getServerFilePath()); err != nil {
		if os.IsNotExist(err) {
			var servers []Server
			WriteAll(servers)
		} else {
			kit.Check(err)
		}
	}
}

func ReadAll() []Server {
	var servers []Server
	file, err := os.Open(getServerFilePath())
	kit.Check(err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&servers)
	return servers
}

func WriteAll(servers []Server) {
	bs, err := json.Marshal(servers)
	kit.Check(err)

	var out bytes.Buffer
	err = json.Indent(&out, bs, "", "\t")
	kit.Check(err)

	file, err := os.Create(getServerFilePath())
	defer file.Close()
	out.WriteTo(file)
}

func Add(sv Server) {
	servers := ReadAll()
	servers = append(servers, sv)

	WriteAll(servers)
}

func FindByName(serverName string) (*Server, bool) {
	servers := ReadAll()
	for _, vs := range servers {
		if vs.ServerName == serverName {
			return &vs, true
		}
	}

	return nil, false

}

func RemoveByName(serverName string) bool {
	servers := ReadAll()
	for i, sv := range servers {
		if sv.ServerName == serverName {
			servers = append(servers[:i], servers[i+1:]...)
			WriteAll(servers)
			return true
		}
	}

	return false
}

func getServerFilePath() string {
	user, err := user.Current()
	kit.Check(err)
	return user.HomeDir + "/" + serverFileName
}

package main

import (
	"bytes"
	"encoding/json"
	"os"
	"os/user"
)

func initServerFile() {
	if _, err := os.Stat(getServerFilePath()); err != nil {
		if os.IsNotExist(err) {
			var servers []Server
			writeAllServer(servers)
		} else {
			check(err)
		}
	}
}

func readAllServer() []Server {
	var servers []Server
	file, err := os.Open(getServerFilePath())
	check(err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&servers)
	return servers
}

func writeAllServer(servers []Server) {
	bs, err := json.Marshal(servers)
	check(err)

	var out bytes.Buffer
	err = json.Indent(&out, bs, "", "\t")
	check(err)

	file, err := os.Create(getServerFilePath())
	defer file.Close()
	out.WriteTo(file)
}

func addServer(sv Server) {
	servers := readAllServer()
	servers = append(servers, sv)

	writeAllServer(servers)
}

func findServerByServerName(serverName string) (*Server, bool) {
	servers := readAllServer()
	for _, vs := range servers {
		if vs.ServerName == serverName {
			return &vs, true
		}
	}

	return nil, false

}

func removeServerByServerName(serverName string) bool {
	servers := readAllServer()
	for i, sv := range servers {
		if sv.ServerName == serverName {
			servers = append(servers[:i], servers[i+1:]...)
			writeAllServer(servers)
			return true
		}
	}

	return false
}

func getServerFilePath() string {
	user, err := user.Current()
	check(err)
	return user.HomeDir + "/" + serverFileName
}

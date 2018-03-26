package main

func createHandler(args []string) handler {
	var command string
	var cas []string // command args
	var casl int     // command args length
	if len(args) > 0 {
		command = args[0]
		cas = args[1:]
		casl = len(cas)
	}
	var h handler

	switch {
	case command == "conn" && casl == 1:
		h = &connection{
			serverName: cas[0],
		}
	case command == "add" && casl == 5:
		h = &add{
			serverName: cas[0],
			username:   cas[1],
			ip:         cas[2],
			port:       cas[3],
			password:   cas[4],
		}
	case command == "rm" && casl == 1:
		h = &remove{
			serverName: cas[0],
		}
	case command == "ls":
		h = &list{}
	case command == "version":
		fallthrough
	case command == "-version":
		fallthrough
	case command == "-v":
		h = &version{}
	default:
		h = &help{}
	}

	return h
}

package command

func CreateHandler(command string) Handler {
	var h Handler

	switch {
	case command == "conn":
		h = &Connect{}
	case command == "add":
		h = &Add{}
	case command == "rm":
		h = &Remove{}
	case command == "ls":
		h = &List{}
	case command == "-version":
		fallthrough
	case command == "-v":
		h = &Version{}
	default:
		h = &Help{}
	}

	return h
}

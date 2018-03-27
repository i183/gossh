package command

type Handler interface {
	Init(args []string) bool //Initialization Handler
	Execute()                //Execute
	Help()                   //Help
}

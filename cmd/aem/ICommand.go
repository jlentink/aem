package main

/*
ICommand interface for all commands that can be invoked
*/
type ICommand interface {
	Init()
	Execute(args []string)
	getOpt(args []string)
	readConfig() bool
	GetCommand() []string
	GetHelp() string
}

package main

import "fmt"

type Alpha struct{}

func (a *Alpha) Execute(args []string) error {
	fmt.Println("Running the alpha command")
	return nil
}

var PluginCommandHook Alpha
var PluginCommandName = "alpha"

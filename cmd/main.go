package main

import (
	"fmt"
	"plugin"

	goflags "github.com/jessevdk/go-flags"
)

func main() {
	parser := goflags.NewNamedParser("cmd", 0)

	addBuiltinCommand(parser)
	addPluginCommand(parser)

	_, err := parser.Parse()
	if err != nil {
		panic(err)
	}
}

func addBuiltinCommand(parser *goflags.Parser) {
	var builtinOps builtIn
	_, err := parser.AddCommand("builtin", "built-in command", "long desc", &builtinOps)
	if err != nil {
		panic(err)
	}
}

func addPluginCommand(parser *goflags.Parser) {
	p, err := plugin.Open("../plug-alpha/plug-alpha.so")
	if err != nil {
		panic(err)
	}

	hook, err := p.Lookup("PluginCommandHook")
	if err != nil {
		panic(err)
	}

	name, err := p.Lookup("PluginCommandName")
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand(*name.(*string), "plugin", "plugin", hook)
	if err != nil {
		panic(err)
	}
}

type builtIn struct{}

func (b *builtIn) Execute(args []string) error {
	fmt.Println("built-in command")
	return nil
}

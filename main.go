package main

import (
	"fmt"
	"os"

	"github.com/ajswetz/go-gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}

	cliState := state{
		configuration: cfg,
	}

	cliCommands := commands{
		commands: make(map[string]func(*state, command) error),
	}

	cliCommands.register("login", handlerLogin)

	cliArguments := os.Args

	if len(cliArguments) < 2 {
		fmt.Println("gator requires a command name, but one was not provided")
		os.Exit(1)
	}

	cmdName := cliArguments[1]
	cmdArgs := cliArguments[2:]
	cmdToRun := command{
		name:      cmdName,
		arguments: cmdArgs,
	}

	err = cliCommands.run(&cliState, cmdToRun)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
	}
}

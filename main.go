package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ajswetz/go-gator/internal/config"
	"github.com/ajswetz/go-gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// Get configuration from ".gatorconfig.json" file
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("An error occurred attempting to read configuration: %v\n", err)
	}

	// Open connection to database
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("An error occurred attempting to connect to postgres db: %v\n", err)
	}

	// Create new *database.Queries object
	dbQueries := database.New(db)

	// Add configuration and db connection to the `state` struct
	programState := state{
		configuration: cfg,
		db:            dbQueries,
	}

	// create new `commands` struct that will hold a map of all available cli commands
	cliCommands := commands{
		commands: make(map[string]func(*state, command) error),
	}

	// register available commands
	cliCommands.register("login", handlerLogin)
	cliCommands.register("register", handlerRegister)

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

	err = cliCommands.run(&programState, cmdToRun)
	if err != nil {
		fmt.Printf("An error occured attempting to run %s: %v\n", cmdName, err)
	}
}

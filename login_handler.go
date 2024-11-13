package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	// Check to see if a `username` argument was passed to the command
	if len(cmd.arguments) == 0 {
		fmt.Println("login handler expects one argument; none were provided")
		os.Exit(1)
	}
	// Set `username` variable to the name passed in as the cmd argument
	username := cmd.arguments[0]

	// Confirm whether given user exists in the database
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Printf("Unable to get user '%s' from the database: %v\n", username, err)
		os.Exit(1)
	}

	// Set current user in the config to the user given to the `login` command
	err = s.configuration.SetUser(username)
	if err != nil {
		return err
	}

	// Print results to the console
	fmt.Printf("%s is now logged in\n", username)

	return nil
}

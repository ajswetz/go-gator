package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ajswetz/go-gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	// Check to see if a `username` argument was passed to the command
	if len(cmd.arguments) == 0 {
		fmt.Println("register handler expects one argument; none were provided")
		os.Exit(1)
	}

	// Set `username` variable to the name passed in as the cmd argument
	username := cmd.arguments[0]

	// Build CreateUserParams struct object to pass to the s.db.CreateUser function
	createUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	// Attempt to create the new user in the database
	newUser, err := s.db.CreateUser(context.Background(), createUserParams)
	if err != nil {
		fmt.Printf("An error occurred attempting to create user in the DB: %v\n", err)
		os.Exit(1)
	}

	// Set current user in the config to the user given to the `register` command
	s.configuration.SetUser(username)
	if err != nil {
		return err
	}

	// print message confirming user was created
	fmt.Printf("New user '%s' was successfully created\n", username)
	fmt.Println("User details:")
	jsonData, err := json.MarshalIndent(newUser, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling data to JSON: %v\n", err)
		fmt.Printf("Raw user data: %+v", newUser)
	} else {
		fmt.Println(string(jsonData))
	}

	return nil
}

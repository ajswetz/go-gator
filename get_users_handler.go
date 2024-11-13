package main

import (
	"context"
	"fmt"
	"os"
)

func handlerUsers(s *state, _ command) error {

	// Attempt to get all users from the database
	allUsers, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		fmt.Printf("An error occurred attempting to get all users from the database: %v\n", err)
		os.Exit(1)
	}

	// Get current user from program state
	currentUser := s.configuration.CurrentUserName

	// Print list of users in the database. Mark which user is the current user
	for _, user := range allUsers {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}

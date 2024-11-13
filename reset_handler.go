package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, _ command) error {

	// Attempt to delete all users in the database
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Printf("An error occurred attempting to delete all users: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("All users successfully deleted from the database")
	}

	return nil

}

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ajswetz/go-gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	// Check to see if `name` and `url` arguments were passed to the command
	if len(cmd.arguments) != 2 {
		fmt.Println("addfeed expects exactly two arguments: the 'name' and 'url' of the feed")
		os.Exit(1)
	}
	// Set name and url variables passed in as arguments
	feedName := cmd.arguments[0]
	feedURL := cmd.arguments[1]

	// Get current user details from the database
	currentUser, err := s.db.GetUser(context.Background(), s.configuration.CurrentUserName)
	if err != nil {
		fmt.Printf("Unable to get user '%s' from the database: %v\n", s.configuration.CurrentUserName, err)
		os.Exit(1)
	}

	// Build database.CreateFeedParams object
	newFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    currentUser.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), newFeedParams)
	if err != nil {
		fmt.Printf("Unable to add feed '%s' with URL '%s'\n", feedName, feedURL)
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Print results to the console
	fmt.Println("Successfully added new feed to the database")
	fmt.Printf("%+v\n", feed)

	return nil
}

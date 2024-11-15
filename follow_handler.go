package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ajswetz/go-gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {

	// Check to see if a `url` argument was passed to the command
	if len(cmd.arguments) != 1 {
		fmt.Println("follow command expects exactly one argument: the url of the feed to follow")
		os.Exit(1)
	}

	// Set `feedURL` variable to the url passed in as the cmd argument
	feedURL := cmd.arguments[0]

	// Build CreateFeedFollowParams object to pass to the CreateFeedFollow() function
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      user.Name,
		Url:       feedURL,
	}

	// Attempt to create the feed follows record
	feedFollowResp, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		fmt.Println("An error occurred when attempting to create a new feed follow record")
		fmt.Printf("Error: %v\n", err)
	}

	// Print results to the console
	fmt.Printf("%s is now following feed '%s'\n", feedFollowResp.UserName, feedFollowResp.FeedName)

	return nil
}

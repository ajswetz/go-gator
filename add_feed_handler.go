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

func handlerAddFeed(s *state, cmd command, user database.User) error {
	// Check to see if `name` and `url` arguments were passed to the command
	if len(cmd.arguments) != 2 {
		fmt.Println("addfeed expects exactly two arguments: the 'name' and 'url' of the feed")
		os.Exit(1)
	}
	// Set name and url variables passed in as arguments
	feedName := cmd.arguments[0]
	feedURL := cmd.arguments[1]

	// Build database.CreateFeedParams object
	newFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), newFeedParams)
	if err != nil {
		fmt.Printf("Unable to add feed '%s' with URL '%s'\n", feedName, feedURL)
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Print results to the console
	fmt.Println("Successfully added new feed to the database")
	jsonData, err := json.MarshalIndent(feed, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling data to JSON: %v\n", err)
		fmt.Printf("Raw feed data: %+v", feed)
	} else {
		fmt.Println(string(jsonData))
	}

	// Create a feed-follow record for the current user and the added feed
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

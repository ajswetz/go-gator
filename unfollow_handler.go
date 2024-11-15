package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ajswetz/go-gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	// Check to see if a `url` argument was passed to the command
	if len(cmd.arguments) != 1 {
		fmt.Println("follow command expects exactly one argument: the url of the feed to follow")
		os.Exit(1)
	}

	// Set `feedURL` variable to the url passed in as the cmd argument
	feedURL := cmd.arguments[0]

	// Build DeleteFeedFollowParams to pass to DeleteFeedFollow() function
	delFeedFollowParams := database.DeleteFeedFollowParams{
		Name: user.Name,
		Url:  feedURL,
	}

	// Attempt to delete the feed follows record
	err := s.db.DeleteFeedFollow(context.Background(), delFeedFollowParams)
	if err != nil {
		fmt.Println("An error occurred when attempting to delete feed follow record")
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Print results to the console
	fmt.Printf("%s has successfully unfollowed feed '%s'\n", user.Name, feedURL)

	return nil
}

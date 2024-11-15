package main

import (
	"context"
	"fmt"

	"github.com/ajswetz/go-gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	// Attempt to get all followed feeds for the currently logged in user
	feedsFollowed, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		fmt.Printf("An error occured attempting to get all followed feeds for %s\n", user.Name)
		fmt.Printf("Error: %v\n", err)
	}

	// Print results to the console
	fmt.Printf("%s is currently following these feeds:\n", user.Name)
	for _, feed := range feedsFollowed {
		fmt.Printf(" - %s\n", feed.FeedName)
	}

	return nil
}

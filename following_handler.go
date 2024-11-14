package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {

	// Attempt to get all followed feeds for the currently logged in user
	feedsFollowed, err := s.db.GetFeedFollowsForUser(context.Background(), s.configuration.CurrentUserName)
	if err != nil {
		fmt.Printf("An error occured attempting to get all followed feeds for %s\n", s.configuration.CurrentUserName)
		fmt.Printf("Error: %v\n", err)
	}

	// Print results to the console
	fmt.Printf("%s is currently following these feeds:\n", s.configuration.CurrentUserName)
	for _, feed := range feedsFollowed {
		fmt.Printf(" - %s\n", feed.FeedName)
	}

	return nil
}

package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, c command) error {

	// Add an agg command. Later this will be our long-running aggregator service.
	// For now, we'll just use it to fetch a single feed and ensure our parsing works.
	// It should fetch the feed found at https://www.wagslane.dev/index.xml and print the entire struct to the console.

	const feedURL = "https://www.wagslane.dev/index.xml"

	fetchedFeed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	} else {
		fmt.Println("Feed fetched:")
		fmt.Printf("%+v", fetchedFeed)
	}

	return nil
}

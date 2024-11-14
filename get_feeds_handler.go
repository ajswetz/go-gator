package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

func handlerGetFeeds(s *state, cmd command) error {

	// Attempt to get all feeds from the database
	allFeeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		fmt.Printf("An error occurred attempting to get all feeds from the database: %v\n", err)
		os.Exit(1)
	}

	// Marshal feed data to JSON and print to the console
	jsonData, err := json.MarshalIndent(allFeeds, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling data to JSON: %v\n", err)
		fmt.Printf("Raw feed data: %+v", allFeeds)
	} else {
		fmt.Println(string(jsonData))
	}

	return nil
}

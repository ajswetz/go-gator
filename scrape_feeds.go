package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ajswetz/go-gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func scrapeFeeds(s *state) {

	// Get the next feed to fetch from the DB.
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("There was a problem getting the next feed to fetch")
		fmt.Printf("Error: %v\n", err)
	}

	// Mark it as fetched.
	markFetchedParams := database.MarkFeedFetchedParams{
		ID:        nextFeed.ID,
		UpdatedAt: time.Now(),
	}
	err = s.db.MarkFeedFetched(context.Background(), markFetchedParams)
	if err != nil {
		fmt.Printf("There was a problem marking feed '%s' as fetched\n", nextFeed.Name)
		fmt.Printf("Error: %v\n", err)
	}

	// Fetch the feed using the URL (we already wrote this function)
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Printf("There was a problem fetching feed '%s' at %s\n", nextFeed.Name, nextFeed.Url)
		fmt.Printf("Error: %v\n", err)
	}

	// Iterate over the items in the feed and save them to the `posts` table in the database
	for _, item := range rssFeed.Channel.Item {

		// fmt.Printf("Adding post '%s' to the database...\n", item.Title)

		const layout1 = "Mon, 02 Jan 2006 15:04:05 -0700"
		const layout2 = "Mon, 02 Jan 2006 15:04:05 MST"

		pubDateTime, err := time.Parse(layout1, item.PubDate)
		if err != nil {
			// Try another layout
			pubDateTime, err = time.Parse(layout2, item.PubDate)
			if err != nil {
				fmt.Println("Unable to parse PubDate string into time.Time value")
				fmt.Println("Zero value will be used for PublishedAt value")
			}
		}

		createPostPararms := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDateTime,
			FeedID:      nextFeed.ID,
		}

		err = s.db.CreatePost(context.Background(), createPostPararms)
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			continue
		}

		if err != nil {
			fmt.Println("Unable to save post to the database")
			fmt.Printf("Error: %v\n", err)
		}

	}
}

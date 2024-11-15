package main

import (
	"fmt"
	"os"
	"time"
)

func handlerAgg(s *state, c command) error {

	// Check to see if a `time_between_reqs` argument was passed to the command
	if len(c.arguments) != 1 {
		fmt.Println("agg handler expects exactly one argument: a time_between_reqs duration string")
		os.Exit(1)
	}

	// Convert time_between_requests string argument to time.Duration value
	time_btw_reqs_str := c.arguments[0]
	timeBetweenRequests, err := time.ParseDuration(time_btw_reqs_str)
	if err != nil {
		fmt.Printf("Unable to convert '%s' to time.Duration value\n", time_btw_reqs_str)
	}

	// Start scraping feeds on infinite loop once every `time_between_requests`
	fmt.Printf("Collecting feeds every %v...\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Println("Running scrapeFeeds()...")
		scrapeFeeds(s)
		fmt.Println()
	}

	return nil
}

package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	// Build HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	// Create new HTTP client
	client := http.Client{}

	// Use HTTP client to execute HTTP request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read HTTP response into `data` variable
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal XML into fetchedFeed variable of type `RSSFeed`
	var fetchedFeed RSSFeed
	if err = xml.Unmarshal(data, &fetchedFeed); err != nil {
		return nil, err
	}

	// Use the html.UnescapeString function to decode escaped HTML entities (like &ldquo;).
	// You'll need to run the Title and Description fields (of both the entire channel
	// as well as the items) through this function.
	fetchedFeed.Channel.Title = html.UnescapeString(fetchedFeed.Channel.Title)
	fetchedFeed.Channel.Description = html.UnescapeString(fetchedFeed.Channel.Description)
	for i := range fetchedFeed.Channel.Item {
		fetchedFeed.Channel.Item[i].Title = html.UnescapeString(fetchedFeed.Channel.Item[i].Title)
		fetchedFeed.Channel.Item[i].Description = html.UnescapeString(fetchedFeed.Channel.Item[i].Description)
	}

	return &fetchedFeed, nil

}

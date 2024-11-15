package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/ajswetz/go-gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	// Check to see if a `limit` argument was passed to the command
	// limit will be set to '2' as the default in case no argument is passed
	limit := 2
	var err error
	if len(cmd.arguments) > 0 {
		limit, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			fmt.Println("It looks like you tried to include a limit with the 'browse' command.")
			fmt.Printf("However, gator was unable to convert '%s' to an integer.", cmd.arguments[0])
			fmt.Println("Gator will exit now. Please try again with valid syntax.")
		}
	}

	// Attempt to get posts for the currently logged in user's feeds
	getPostsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), getPostsParams)
	if err != nil {
		fmt.Printf("An error occured getting posts for user '%s'\n", user.Name)
		os.Exit(1)
	}

	// Print results to the console
	fmt.Println("Here are your most recent posts:")
	fmt.Println()
	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Description: %s\n", post.Description)
		fmt.Printf("Publication Date: %v\n", post.PublishedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println()
	}
	return nil
}

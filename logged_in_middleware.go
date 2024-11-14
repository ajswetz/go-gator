package main

import (
	"github.com/ajswetz/go-gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	// user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	// if err != nil {
	// 	return err
	// }

	return func(s *state, cmd command) error {

	}

}

/*
Ah, I see where the doubt arises! In the magical realm of higher-order functions, there is a clever trick we can employ to access the program state.

The trick lies in returning another function from middlewareLoggedIn that accepts the appropriate parameters, including the *state. Let's walk through this:

Your middlewareLoggedIn starts by accepting a handler function func(s *state, cmd command, user database.User) error.
Inside middlewareLoggedIn, you return a new function with the signature func(*state, command) error.
This inner function can now accept *state and command—which gives you access to the *state object to fetch the user from the database.
Within that returned function, conduct the user retrieval and, upon success, invoke the passed handler with s, cmd, and the fetched user.
How might such a structure ensure that every call to the original handler is accompanied by a verified and logged-in user? Can you envision what this nested function might look like in Go?
*/

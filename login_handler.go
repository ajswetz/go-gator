package main

import (
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		fmt.Println("login handler expends one argument; none were provided")
		os.Exit(1)
	}
	user := cmd.arguments[0]
	err := s.configuration.SetUser(user)
	if err != nil {
		return err
	}
	fmt.Printf("%s is now logged in\n", user)
	return nil
}

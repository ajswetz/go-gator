package main

type command struct {
	name      string
	arguments []string
}

// This will hold all the commands the CLI can handle
type commands struct {
	commands map[string]func(*state, command) error // Map of command names to their handler functions.
}

// This method registers a new handler function for a command name.
func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

// This method runs a given command with the provided state if it exists.
func (c *commands) run(s *state, cmd command) error {
	err := c.commands[cmd.name](s, cmd)
	if err != nil {
		return err
	} else {
		return nil
	}
}

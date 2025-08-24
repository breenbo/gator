package internal

import (
	"fmt"

	"github.com/breenbo/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no argument")
	}

	// for login, only one argument
	s.cfg.SetUser(cmd.arguments[0])

	fmt.Printf("User %v set\n", cmd.arguments[0])
	return nil
}

type commands struct{}

package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/breenbo/gator/internal/config"
	"github.com/breenbo/gator/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *config.Config
}

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Commands map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	if f, ok := c.Commands[cmd.Name]; !ok {
		return fmt.Errorf("command not found")
	} else {
		if err := f(s, cmd); err != nil {
			return err
		}
	}

	return nil
}

func (c *Commands) Register(s string, f func(*State, Command) error) {
	c.Commands[s] = f
}

func RegisterFn() Commands {
	c := Commands{
		Commands: make(map[string]func(*State, Command) error, 10),
	}

	c.Register("login", HandlerLogin)
	c.Register("register", HandlerRegister)
	c.Register("reset", HandleReset)
	c.Register("users", HandleList)
	c.Register("agg", HandleAgg)

	return c
}

func GetUserEntries() Command {
	// get command and args from cli
	args := os.Args

	if len(args) < 2 {
		log.Fatal("no enough arguments\n")
	}

	return Command{
		Name:      args[1],
		Arguments: args[2:],
	}
}

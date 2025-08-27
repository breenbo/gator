package internal

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/breenbo/gator/internal/config"
	"github.com/breenbo/gator/internal/database"
	"github.com/google/uuid"
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
	fmt.Printf("running a command: %s\n", cmd.Name)

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

	return c
}

func GetUserEntries() Command {
	// get command and args from cli
	args := os.Args

	if len(args) < 2 {
		fmt.Print("no enough arguments\n")
		os.Exit(1)
	}

	return Command{
		Name:      args[1],
		Arguments: args[2:],
	}
}

//
//
//

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("username needed")
	}

	ctx := context.Background()
	name := cmd.Arguments[0]
	if _, err := s.Db.GetUser(ctx, name); err != nil {
		fmt.Print("not in db\n")
		os.Exit(1)
	}

	// for login, only one argument
	s.Cfg.SetUser(name)
	fmt.Printf("User %v set\n", name)

	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("name needed")
	}

	ctx := context.Background()
	name := cmd.Arguments[0]
	// if user already in db, return exit(1)
	if _, err := s.Db.GetUser(ctx, name); err == nil {
		fmt.Printf("%s already in db\n", name)
		os.Exit(1)
	}

	// create user in db
	uuid := uuid.New()
	now := time.Now()

	newUser := database.CreateUserParams{
		ID:        uuid.String(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	}

	if _, err := s.Db.CreateUser(ctx, newUser); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set in the config file
	s.Cfg.SetUser(name)
	fmt.Printf("User %s created\n", name)

	return nil
}

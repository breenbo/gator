package internal

import (
	"context"

	"github.com/breenbo/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(s *State, cmd Command) error {

	fn := func(s *State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.Current_user_name)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}

	return fn
}

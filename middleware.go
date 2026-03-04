package main

import (
	"context"
	"github.com/benbunsford/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		err = handler(s, cmd, user)
		if err != nil {
			return err
		}
		return nil
	}
}

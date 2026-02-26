package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("The login command expects a username. Enter your username after 'login'.")
	}

	//checks if user with provided name exists
	_, err := s.db.GetUserByName(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Username has been updated. Current user: %v", cmd.args[0])

	return nil
}

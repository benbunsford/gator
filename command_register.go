package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/benbunsford/gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("The register command expects a name. Enter a name after 'register' to add a new user.")
	}

	userData := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	_, err := s.db.CreateUser(context.Background(), userData)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(cmd.args[0])

	fmt.Printf("User was created: %v\n", userData)

	return nil
}

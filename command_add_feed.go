package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/benbunsford/gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("The addfeed command expects a name and a url. Example: addfeed 'Big News RSS' 'https://bignews.org/new'")
	}

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedData := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedData)
	if err != nil {
		return err
	}

	fmt.Printf("feed added: %v\n", feed)

	return nil
}

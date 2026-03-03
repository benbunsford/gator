package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/benbunsford/gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func followHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("The follow command expects a url. Example: 'follow www.nbc.com/new")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	followData := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), followData)
	if err != nil {
		return nil
	}

	fmt.Printf("%v now follows %v", s.cfg.CurrentUserName, feed.Name)

	return nil
}

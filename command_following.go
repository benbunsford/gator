package main

import (
	"context"
	"fmt"
	"github.com/benbunsford/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("%v is following:\n", user.Name)
	for _, follow := range follows {
		fmt.Printf("%v\n", follow.FeedName)
	}

	return nil
}

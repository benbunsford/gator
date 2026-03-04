package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

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

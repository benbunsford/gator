package main

import (
	"context"
	"fmt"

	"github.com/benbunsford/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("unfollow expects the url as an arg. example: unfollow 'www.bignews.com'")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	deleteData := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), deleteData)
	if err != nil {
		return err
	}

	fmt.Printf("%v has unfollowed %v\n", user.Name, feed.Name)

	return nil
}

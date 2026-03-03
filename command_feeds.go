package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feedList, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feedList {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("Name: %v, URL: %v, User: %v\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}

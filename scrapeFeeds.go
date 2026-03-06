package main

import (
	"context"
	"fmt"
	"time"

	"github.com/benbunsford/gator/internal/database"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	markFeedData := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        nextFeed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), markFeedData)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for i := 0; i < len(rssFeed.Channel.Item); i++ {
		fmt.Printf("%v\n", rssFeed.Channel.Item[i].Title)
	}

	return nil
}

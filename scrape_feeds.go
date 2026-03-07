package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/benbunsford/gator/internal/database"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
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

	fmt.Printf("Scraped %v posts\n", len(rssFeed.Channel.Item))

	for _, item := range rssFeed.Channel.Item {
		formats := []string{time.RFC1123Z, time.RFC1123, time.RFC3339}
		var publishedAt time.Time
		for _, format := range formats {
			publishedAt, err = time.Parse(format, item.PubDate)
			if err != nil {
				log.Printf("could not parse published_at: %v\n", item.PubDate)
			} else {
				break
			}
		}

		postData := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: sql.NullTime{
				Time:  publishedAt,
				Valid: !publishedAt.IsZero(),
			},
			FeedID: nextFeed.ID,
		}

		_, err := s.db.CreatePost(context.Background(), postData)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("error creating post: %v\n", err)
		}
	}

	return nil
}

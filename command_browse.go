package main

import (
	"context"
	"fmt"
	"github.com/benbunsford/gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.args) < 1 {
		fmt.Println("hint: you can specify the number of posts you want to see after the browse command. Default is 2. Example: 'browse 4'")
	} else {
		limitInt, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(limitInt)
	}

	postsData := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), postsData)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf(`
%v - %v
%v

%v

Read here: %v


-------------------------------------------------------------
			`, post.FeedName, post.Title, post.PublishedAt.Time, post.Description.String, post.Url)
	}

	return nil
}

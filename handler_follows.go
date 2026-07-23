package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/bubskee/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	tStamp := time.Now().UTC()
	user_name := s.cfg.CurrentUserName

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error while fetching feed: %w", err)
	}

	user, err := s.db.GetUser(context.Background(), user_name)
	if err != nil {
		return fmt.Errorf("error while fetching user: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(
		context.Background(), database.CreateFeedFollowParams{
			ID: uuid.New(),
			CreatedAt: tStamp,
			UpdatedAt: tStamp,
			FeedID: feed.ID,
			UserID: user.ID,
		})
	if err != nil {
		return fmt.Errorf("error while creating feedfollow: %w", err)
	}

	fmt.Println("Feed successfully followed!")
	fmt.Printf("\tfeed: %v\n", feedFollow.FeedName)
	fmt.Printf("\tuser: %v\n", feedFollow.UserName)

	return nil
}
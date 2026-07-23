package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/bubskee/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	tStamp := time.Now().UTC()

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error while fetching feed: %w", err)
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

	fmt.Println("Feed successfully followed:")
	fmt.Printf(" * feed:\t %v\n", feedFollow.FeedName)
	fmt.Printf(" * user:\t %v\n", feedFollow.UserName)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error while fetching follows: %w", err)
	}

	numFollows := len(follows)
	if numFollows < 1 {
		fmt.Printf("User: %v\n", user.Name)
		fmt.Println("Has no followed RSS feeds!")
		return nil
	}

	fmt.Printf("User: %v\n", user.Name)
	fmt.Printf("Found %v followed RSS feeds:\n", numFollows)
	for i, f := range follows {
		fmt.Printf("-- %v ---------------\n", i+1)
		fmt.Printf(" * feed name:\t %v\n", f.FeedName)
	}
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	err := s.db.DeleteFollowByUrl(
		context.Background(), database.DeleteFollowByUrlParams{
		 	Url: url,
		 	UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error while deleting feedfollow: %w", err)
	}

	fmt.Printf(" * feed url:\t %v\n", url)
	fmt.Printf("Successfully unfollowed for user: %v\n", user.Name)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

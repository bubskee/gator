package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeedsWithUserName(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) < 1 {
		fmt.Println("No Saved RSS feeds!")
		return nil
	}

	fmt.Println("Saved RSS feeds:")
	for i, f := range feeds {
		fmt.Printf("-- %v ---------------\n", i+1)
		fmt.Printf(" * feed name:\t %v\n", f.Name)
		fmt.Printf(" * feed url:\t %v\n", f.Url)
		fmt.Printf(" * saved by:\t %v\n", f.User)
	}
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

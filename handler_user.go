package main

import (
	"fmt"
	"time"
	"context"

	"github.com/google/uuid"

	"github.com/bubskee/gator/internal/database")

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	tStamp := time.Now().UTC()

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: tStamp,
		UpdatedAt: tStamp,
		Name: name,
	})
	if err != nil {
		return fmt.Errorf("couldn't register user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User registered successfully!")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user with name: %s \nerror: %w", name, err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	
	err := s.db.ClearUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error while resetting users table: %v", err)
	}

	fmt.Println("Users table reset!")
	return nil
}

func printUser(user database.User) {
	fmt.Println(" -- User Data: ")
	fmt.Printf("        -- Name: %s\n", user.Name)
	fmt.Printf("        -- UUID: %v\n", user.ID)
}

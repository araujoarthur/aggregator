package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/araujoarthur/aggregator/internal/database"
	"github.com/google/uuid"
)

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	usr := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      cmd.Args[0],
	}

	if s.UserExists(usr.Name) {
		return fmt.Errorf("name already exists")
	}

	if _, err := s.DbQueries.CreateUser(context.Background(), usr); err != nil {
		return fmt.Errorf("something went wrong during user creation: %w", err)
	}

	err := s.RegisteredCommands.run(s, command{Name: "login", Args: cmd.Args})
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created and switched successfully!")
	return nil
}

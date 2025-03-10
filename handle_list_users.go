package main

import (
	"context"
	"fmt"
)

func handleListUsers(s *state, cmd command) error {
	users, err := s.DbQueries.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("something went wrong returning users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s \n", user.Name)
	}

	return nil
}

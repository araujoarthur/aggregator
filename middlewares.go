package main

import (
	"context"
	"fmt"

	"github.com/araujoarthur/aggregator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *state,
		cmd command,
		user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		usr, err := s.DbQueries.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("failed to fetch user data: %w (are you logged in?)", err)
		}
		return handler(s, cmd, usr)
	}
}

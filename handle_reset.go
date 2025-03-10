package main

import (
	"context"
	"fmt"
)

func handleReset(s *state, cmd command) error {
	if err := s.DbQueries.ResetUsers(context.Background()); err != nil {
		return err
	}

	fmt.Println("Reset successful!")
	return nil
}

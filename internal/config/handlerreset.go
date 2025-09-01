package config

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.DB.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset users")
	}
	fmt.Printf("users are reset\n")
	return nil
}

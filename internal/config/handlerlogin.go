package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required")
	}

	username := cmd.Args[0]

	user, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("error: user %q does not exist\n", username)
			os.Exit(1)
		}
		return fmt.Errorf("failed to fetch user %q: %w", username, err)
	}

	if err := s.Config.SetUser(username); err != nil {
		return fmt.Errorf("failed to save user in config: %w", err)
	}

	fmt.Printf("user %q has been set successfully\n", username)
	fmt.Printf("debug: logged-in user row = %+v\n", user)
	return nil
}

package config

import (
	"context"
	"fmt"
	"time"

	"github.com/Darkpowercross/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]

	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return fmt.Errorf("user %q already exists", username)
		}
		return fmt.Errorf("failed to insert user: %w", err)
	}

	if err := s.Config.SetUser(username); err != nil {
		return fmt.Errorf("failed to update config: %w", err)
	}

	fmt.Printf("user %q saved to database and set in config\n", username)
	fmt.Printf("debug: created user row = %+v\n", user)
	return nil
}

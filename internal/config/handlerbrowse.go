package config

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Darkpowercross/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2

	if len(cmd.Args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specifiedLimit
			fmt.Printf("limit is %d", limit)
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.DB.GetPosts(context.Background(), database.GetPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %v\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %v ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %v\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}

package config

import (
	"fmt"

	"github.com/Darkpowercross/gator/internal/database"
)

type State struct {
	DB     *database.Queries
	Config *Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	handlers map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	if len(cmd.Name) == 0 {
		return fmt.Errorf("command name is required")
	}

	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	return handler(s, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if len(name) == 0 || f == nil {
		return
	}

	if c.handlers == nil {
		c.handlers = make(map[string]func(*State, Command) error)
	}

	c.handlers[name] = f
}

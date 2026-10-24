package main

import (
	"errors"
	"fmt"
	"gator/internal/config"
)

type state struct{ Config *config.Config }

type command struct {
	Name string
	Args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Login requires an username as argument")
	}
	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set as: '%s'", cmd.Args[0])
	return nil
}

type commands struct {
	Handlers map[string](func(*state, command) error)
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return errors.New("command does not exist")
	}

	return handler(s, cmd)
}

func (c *commands) Register(name string, f func(*state, command) error) {
	c.Handlers[name] = f
}

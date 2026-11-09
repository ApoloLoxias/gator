package main

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"github.com/google/uuid"
	"time"
)

type state struct {
	Config *config.Config
	Db     *database.Queries
}

type command struct {
	Name string
	Args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Login requires an username as argument")
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("user %s not found, error: %v", cmd.Args[0], err)
	}

	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set as: '%s'", cmd.Args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("Register requires a name as its only argument")
	}

	name := cmd.Args[0]

	user, err := s.Db.CreateUser(
		context.Background(), database.CreateUserParams{
			ID:         uuid.New(),
			Name:       name,
			CreateadAt: time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to created user %s, error: %v", name, err)
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("Successfull user registration")
	fmt.Println(user)

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

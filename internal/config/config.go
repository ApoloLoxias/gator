package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	object, err := json.Marshal(c)
	if err != nil {
		return err
	}

	configPath, err := configPath()
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, object, 0666)
	return err
}

func Read() (Config, error) {
	var config Config

	configPath, err := configPath()
	if err != nil {
		return config, err
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, err
}

const configFileName string = ".gatorconfig.json"

func configPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, configFileName), err
}

type State struct{ Config *Config }

type Command struct {
	Name string
	Args []string
}

func HandlerLogin(s *State, cmd Command) error {
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

type Commands struct {
	Handlers map[string](func(*State, Command) error)
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return errors.New("command does not exist")
	}

	return handler(s, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}

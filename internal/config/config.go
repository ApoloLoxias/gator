package config

import (
	"encoding/json"
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

package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) error {
	// Writes config struct to JSON file after setting the current_user_name field

	c.CurrentUserName = user
	err := c.write()
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) write() error {

	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		return err
	}

	return nil
}

func Read() (*Config, error) {
	// Reads the JSON file found at ~/.gatorconfig.json and returns a Config struct.
	// It should read the file from the HOME directory, then decode the JSON string into a new Config struct.

	var configuration Config

	filePath, err := getConfigFilePath()
	if err != nil {
		return &configuration, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return &configuration, err
	}

	if err = json.Unmarshal(data, &configuration); err != nil {
		return &configuration, err
	}

	return &configuration, nil

}

func getConfigFilePath() (string, error) {
	basePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := basePath + "/" + configFileName
	return filePath, nil
}

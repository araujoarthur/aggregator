package config

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	data, err := os.ReadFile(getConfigFilePath())
	if err != nil {
		return Config{}, err
	}

	decoder := json.NewDecoder(bytes.NewBuffer((data)))
	var decoded Config
	if err = decoder.Decode(&decoded); err != nil {
		return Config{}, err
	}

	return decoded, nil
}

func (c *Config) SetUser(u string) error {
	c.CurrentUserName = u
	if err := c.write(); err != nil {
		return err
	}

	return nil
}

func (c *Config) write() error {
	f, err := os.OpenFile(getConfigFilePath(), os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)

	if err = encoder.Encode(c); err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() string {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Failed to Get Home Directory")
	}

	return homeDir + "/" + configFileName
}

package config

import (
	"encoding/json"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url            string
	Current_user_name string
}

func (c *Config) SetUser(userName string) {
	c.Current_user_name = userName

	err := write(*c)
	if err != nil {
		log.Fatal("Unable to setup user")
	}
}

func Read() Config {
	config := Config{}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		log.Fatal("No config path available\n")
	}

	filePath := configFilePath + "/" + configFileName
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Unable to read the config file\n")
	}

	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal("Unable to read the config file\n")
	}

	return config
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		log.Fatal("No config path available\n")
	}

	filePath := configFilePath + "/" + configFileName

	// allow only rw for owner, nothing for others
	err = os.WriteFile(filePath, data, 0600)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

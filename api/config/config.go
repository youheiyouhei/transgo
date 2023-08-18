package config

import (
	"os"
	"path/filepath"
)

var configFile string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Could not find user's home directory: " + err.Error())
	}
	configFile = filepath.Join(home, ".transgo")
}

func SetAPIKey(key string) error {
	return os.WriteFile(configFile, []byte(key), 0600)
}

func GetAPIKey() (string, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	data := fmt.Sprintf("api_key=%s", key)
	return os.WriteFile(configFile, []byte(data), 0600)
}

func GetAPIKey() (string, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		keyPair := strings.SplitN(line, "=", 2)
		if len(keyPair) == 2 && keyPair[0] == "api_key" {
			return keyPair[1], nil
		}
	}

	if scanner.Err() != nil {
		return "", scanner.Err()
	}

	return "", fmt.Errorf("api_key not found in config file")
}

package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// FileOperator provides an interface for file operations, making it testable.
type FileOperator interface {
	Open(filename string) (io.ReadCloser, error)
	Write(filename string, data []byte, perm os.FileMode) error
}

// OSFileOperator is a real FileOperator that interacts with the OS.
type OSFileOperator struct{}

func (f OSFileOperator) Open(filename string) (io.ReadCloser, error) {
	return os.Open(filename)
}

func (f OSFileOperator) Write(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

// Default file operator is the OSFileOperator.
var defaultOperator FileOperator = OSFileOperator{}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find user's home directory: %v", err)
	}
	return filepath.Join(home, ".transgo"), nil
}

func SetAPIKey(key string) error {
	return SetAPIKeyWithOperator(key, defaultOperator)
}

func SetAPIKeyWithOperator(key string, op FileOperator) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data := fmt.Sprintf("api_key=%s", key)
	return op.Write(configPath, []byte(data), 0600)
}

func GetAPIKey() (string, error) {
	return GetAPIKeyWithOperator(defaultOperator)
}

func GetAPIKeyWithOperator(op FileOperator) (string, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return "", err
	}
	file, err := op.Open(configPath)
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

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var configFile string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Could not find user's home directory: " + err.Error())
	}
	configFile = filepath.Join(home, ".transgo")
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringP("set", "s", "", "Set a configuration key in the format key=value")
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage application configuration",
	Long: `Allows you to set or get configuration parameters for the application.
For example:
./transgo config --set api_key=YOUR_API_KEY`,
	Run: func(cmd *cobra.Command, args []string) {
		setKeyValue, _ := cmd.Flags().GetString("set")

		if setKeyValue != "" {
			parts := strings.SplitN(setKeyValue, "=", 2)
			if len(parts) != 2 {
				fmt.Println("Error: Configuration should be in the format key=value")
				return
			}

			key := parts[0]
			value := parts[1]

			if key == "api_key" {
				err := setAPIKey(value)
				if err != nil {
					fmt.Println("Error setting API key:", err)
				} else {
					fmt.Println("API key set successfully.")
				}
			} else {
				fmt.Printf("Unknown configuration key: %s\n", key)
			}
		}
	},
}

func setAPIKey(key string) error {
	return os.WriteFile(configFile, []byte(key), 0600)
}

func getAPIKey() (string, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

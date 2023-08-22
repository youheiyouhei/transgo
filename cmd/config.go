package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/youheiyouhei/transgo/api/config"
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
	Run: executeConfigCmd,
}

func executeConfigCmd(cmd *cobra.Command, args []string) {
	setKeyValue, _ := cmd.Flags().GetString("set")

	if setKeyValue != "" {
		err := setConfiguration(setKeyValue)
		if err != nil {
			fmt.Println("Error setting configuration:", err)
			return
		}
		fmt.Println("Configuration set successfully.")
	} else {
		apiKey, err := fetchCurrentConfiguration()
		if err != nil {
			fmt.Println("Error fetching API key:", err)
			return
		}
		fmt.Println(formatConfiguration(apiKey))
	}
}

func setConfiguration(kv string) error {
	key, value, err := parseKeyValue(kv)
	if err != nil {
		return err
	}

	return applyConfiguration(key, value)
}

func parseKeyValue(kv string) (string, string, error) {
	parts := strings.SplitN(kv, "=", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("Configuration should be in the format key=value and neither key nor value should be empty")
	}
	return parts[0], parts[1], nil
}

func applyConfiguration(key, value string) error {
	switch key {
	case "api_key":
		return config.SetAPIKey(value)
	default:
		return fmt.Errorf("Unknown configuration key: %s", key)
	}
}

func fetchCurrentConfiguration() (string, error) {
	return config.GetAPIKey()
}

func formatConfiguration(apiKey string) string {
	return fmt.Sprintf("Current configuration:\napi_key: %s", apiKey)
}

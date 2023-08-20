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
		setConfiguration(setKeyValue)
	} else {
		displayCurrentConfiguration()
	}
}

func setConfiguration(kv string) {
	parts := strings.SplitN(kv, "=", 2)
	if len(parts) != 2 {
		fmt.Println("Error: Configuration should be in the format key=value")
		return
	}

	key, value := parts[0], parts[1]

	switch key {
	case "api_key":
		if err := config.SetAPIKey(value); err != nil {
			fmt.Println("Error setting API key:", err)
		} else {
			fmt.Println("API key set successfully.")
		}
	default:
		fmt.Printf("Unknown configuration key: %s\n", key)
	}
}

func displayCurrentConfiguration() {
	apiKey, err := config.GetAPIKey()
	if err != nil {
		fmt.Println("Error fetching API key:", err)
		return
	}
	fmt.Printf("Current configuration:\napi_key: %s\n", apiKey)
}

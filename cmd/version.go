package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "v0.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Transgo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Transgo version", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

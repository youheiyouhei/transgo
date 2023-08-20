package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/youheiyouhei/transgo/api/deepl"
)

var languagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "Lists available languages for translation",
	Long: `Fetches and lists the languages that are available for translation
using the DeepL API.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := deepl.NewDeeplClient()
		languages, err := client.GetSupportedLanguages()
		if err != nil {
			fmt.Printf("Error fetching languages: %v\n", err)
			return
		}

		fmt.Println("Available languages:")
		for _, lang := range languages {
			fmt.Printf("- %s (%s)\n", lang.Name, lang.Code)
		}
	},
}

func init() {
	rootCmd.AddCommand(languagesCmd)
}

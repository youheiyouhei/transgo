package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/youheiyouhei/transgo/api/deepl"
	"github.com/youheiyouhei/transgo/interfaces"
)

// languagesCmd represents the 'languages' command
var languagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "Lists available languages for translation",
	Long: `Fetches and lists the languages that are available for translation
using the DeepL API.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := deepl.NewDeeplClient()
		executeLanguagesCmd(client)
	},
}

func init() {
	rootCmd.AddCommand(languagesCmd)
}

func executeLanguagesCmd(client *deepl.DeeplClient) {
	languages, err := fetchSupportedLanguages(client)
	if err != nil {
		fmt.Printf("Error fetching languages: %v\n", err)
		return
	}

	fmt.Println(formatSupportedLanguages(languages))
}

func fetchSupportedLanguages(client *deepl.DeeplClient) (interfaces.SupportedLanguages, error) {
	return client.GetSupportedLanguages()
}

func formatSupportedLanguages(languages interfaces.SupportedLanguages) string {
	var output strings.Builder
	output.WriteString("Available languages:\n")
	for _, lang := range languages {
		output.WriteString(fmt.Sprintf("- %s (%s)\n", lang.Name, lang.Code))
	}
	return output.String()
}

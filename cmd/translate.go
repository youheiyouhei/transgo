package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/youheiyouhei/transgo/api/deepl"
	"github.com/youheiyouhei/transgo/interfaces"
)

var translateCmd = &cobra.Command{
	Use:   "translate [text]",
	Short: "Translates text from a source language to a target language",
	Long: `Translates text from a specified source language to a specified target language using an example translation API.
For example:
./appname translate --source=en --target=ja "Hello, world!"`,
	Args: cobra.ExactArgs(1), // Expect exactly one argument: the text to translate
	Run: func(cmd *cobra.Command, args []string) {
		source, _ := cmd.Flags().GetString("source")
		target, _ := cmd.Flags().GetString("target")
		text := args[0] // get the text from the arguments

		translatedText, err := handleTranslation(text, source, target, deepl.NewDeeplClient())
		if err != nil {
			fmt.Println("Translation failed.", err)
			return
		}

		fmt.Println("Translated text:", translatedText)
	},
}

func handleTranslation(text, source, target string, client translator.Translator) (string, error) {
	translatedText, err := client.Translate([]string{text}, source, target)
	if err != nil {
		fmt.Println("Translation failed.", err)
		return "", err
	}

	return translatedText, nil
}

func init() {
	rootCmd.AddCommand(translateCmd)
	translateCmd.Flags().StringP("source", "s", "", "Source language")
	translateCmd.Flags().StringP("target", "t", "en", "Target language")
}

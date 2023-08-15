/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate [text]",
	Short: "Translates text from a source language to a target language",
	Long: `Translates text from a specified source language to a specified target language using an example translation API.
For example:
./appname translate --source=en --target=ja "Hello, world!"`,
	Args: cobra.ExactArgs(1),  // Expect exactly one argument: the text to translate
	Run: func(cmd *cobra.Command, args []string) {
		source, _ := cmd.Flags().GetString("source")
		target, _ := cmd.Flags().GetString("target")
		text := args[0]  // get the text from the arguments

		translatedText := dummyTranslate(text, source, target)

		fmt.Println("Translated text:", translatedText)
	},
}

func dummyTranslate(text string, source string, target string) string {
	// This is a dummy function. In real-world scenario, you'd call an actual translation API here.
	return fmt.Sprintf("[%s -> %s] %s", source, target, text)
}

func init() {
	rootCmd.AddCommand(translateCmd)

	translateCmd.Flags().StringP("source", "s", "en", "Source language")
	translateCmd.Flags().StringP("target", "t", "ja", "Target language")
}

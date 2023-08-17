/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"bytes"
	"encoding/json"
	"net/http"
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

		translatedText, err := translateWithDeepl(text, source, target)
		if err != nil {
			fmt.Println("Translation failed.", err)
			return
		}

		fmt.Println("Translated text:", translatedText)
	},
}

const deeplAPIEndpoint = "https://api-free.deepl.com/v2/translate"

type DeeplRequest struct {
	Texts   []string `json:"text"`
	Source   string `json:"source_lang"`
	Target   string `json:"target_lang"`
}

type DeeplResponse struct {
	Translations []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

func translateWithDeepl(text, source, target string) (string, error) {
	apiKey, err := getAPIKey()
	if err != nil {
		return "", fmt.Errorf("could not get API key: %v", err)
	}

	requestPayload := DeeplRequest{
		Texts:   []string{text},
		Source:  source,
		Target:  target,
	}

	data, err := json.Marshal(requestPayload)
	if err != nil {
		return "", fmt.Errorf("could not marshal request data: %v", err)
	}

	req, err := http.NewRequest("POST", deeplAPIEndpoint, bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("could not create new request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key " + apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not make request to DeepL: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response body: %v", err)
	}

	var deeplResp DeeplResponse
	err = json.Unmarshal(body, &deeplResp)
	if err != nil {
		return "", fmt.Errorf("could not unmarshal response data: %v", err)
	}

	if len(deeplResp.Translations) == 0 {
		return "", fmt.Errorf("no translations returned by DeepL")
	}

	return deeplResp.Translations[0].Text, nil
}


func init() {
	rootCmd.AddCommand(translateCmd)

	translateCmd.Flags().StringP("source", "s", "en", "Source language")
	translateCmd.Flags().StringP("target", "t", "ja", "Target language")
}

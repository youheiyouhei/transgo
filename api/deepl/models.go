package deepl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/youheiyouhei/transgo/api/config"
	"github.com/youheiyouhei/transgo/translator"
)

type DeeplClient struct{}

func NewDeeplClient() *DeeplClient {
	return &DeeplClient{}
}

func (d *DeeplClient) Translate(request translator.TranslationRequest) (string, error) {
	const deeplAPIEndpoint = "https://api-free.deepl.com/v2/translate"
	data, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("could not marshal request data: %v", err)
	}

	req, err := http.NewRequest("POST", deeplAPIEndpoint, bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("could not create new request: %v", err)
	}
	apiKey, err := config.GetAPIKey()
	if err != nil {
		return "", fmt.Errorf("could not get API key: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)

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

	var deeplResp *translator.TranslationResponse
	err = json.Unmarshal(body, &deeplResp)
	if err != nil {
		return "", fmt.Errorf("could not unmarshal response data: %v", err)
	}

	if len(deeplResp.Translations) == 0 {
		return "", fmt.Errorf("no translations returned by DeepL")
	}

	return deeplResp.Translations[0].Text, nil
}

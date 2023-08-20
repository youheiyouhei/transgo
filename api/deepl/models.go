package deepl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/youheiyouhei/transgo/api/config"
	"github.com/youheiyouhei/transgo/interfaces"
)

const (
	deeplAPIEndpoint       = "https://api-free.deepl.com/v2/translate"
	deeplLanguagesEndpoint = "https://api-free.deepl.com/v2/languages"
)

type DeeplClient struct{}

func NewDeeplClient() *DeeplClient {
	return &DeeplClient{}
}

type TranslationRequest struct {
	Texts  []string `json:"text"`
	Source string   `json:"source_lang"`
	Target string   `json:"target_lang"`
}

type TranslationResponse struct {
	Translations []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

type LanguageResponse struct {
	Language          string `json:"language"`
	Name              string `json:"name"`
	SupportsFormality bool   `json:"supports_formality"`
}

func (d *DeeplClient) Translate(texts []string, source string, target string) (string, error) {
	request := TranslationRequest{
		Texts:  texts,
		Source: source,
		Target: target,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("could not marshal request data: %v", err)
	}

	req, err := d.createRequest("POST", deeplAPIEndpoint, data)
	if err != nil {
		return "", err
	}

	respBody, err := d.sendRequest(req)
	if err != nil {
		return "", err
	}

	return d.parseResponse(respBody)
}

func (d *DeeplClient) GetSupportedLanguages() (interfaces.SupportedLanguages, error) {
	req, err := http.NewRequest("GET", deeplLanguagesEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create new request: %v", err)
	}

	apiKey, err := config.GetAPIKey()
	if err != nil {
		return nil, fmt.Errorf("could not get API key: %v", err)
	}

	req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)

	respBody, err := d.sendRequest(req)
	if err != nil {
		return nil, err
	}

	res, err := d.parseLanguageResponse(respBody)
	if err != nil {
		return nil, err
	}

	var supportedLanguages interfaces.SupportedLanguages
	for _, v := range res {
		supportedLanguages = append(supportedLanguages, interfaces.SupportedLanguage{
			Code: v.Language,
			Name: v.Name,
		})
	}

	return supportedLanguages, nil
}

func (d *DeeplClient) parseLanguageResponse(body []byte) ([]LanguageResponse, error) {
	var languages []LanguageResponse
	if err := json.Unmarshal(body, &languages); err != nil {
		return nil, fmt.Errorf("could not unmarshal language response data: %v", err)
	}

	return languages, nil
}

func (d *DeeplClient) createRequest(method, endpoint string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("could not create new request: %v", err)
	}

	apiKey, err := config.GetAPIKey()
	if err != nil {
		return nil, fmt.Errorf("could not get API key: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)
	return req, nil
}

func (d *DeeplClient) sendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not make request to DeepL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received unexpected status %v from DeepL", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func (d *DeeplClient) parseResponse(body []byte) (string, error) {
	var deeplResp TranslationResponse
	if err := json.Unmarshal(body, &deeplResp); err != nil {
		return "", fmt.Errorf("could not unmarshal response data: %v", err)
	}

	if len(deeplResp.Translations) == 0 {
		return "", fmt.Errorf("no translations returned by DeepL")
	}

	return deeplResp.Translations[0].Text, nil
}

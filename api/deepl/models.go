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

const deeplAPIEndpoint = "https://api-free.deepl.com/v2/translate"

type DeeplClient struct{}

func NewDeeplClient() *DeeplClient {
	return &DeeplClient{}
}

func (d *DeeplClient) Translate(request translator.TranslationRequest) (string, error) {
	data, err := d.marshalRequest(request)
	if err != nil {
		return "", err
	}

	req, err := d.createRequest(data)
	if err != nil {
		return "", err
	}

	respBody, err := d.sendRequest(req)
	if err != nil {
		return "", err
	}

	return d.parseResponse(respBody)
}

func (d *DeeplClient) marshalRequest(request translator.TranslationRequest) ([]byte, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request data: %v", err)
	}
	return data, nil
}

func (d *DeeplClient) createRequest(data []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", deeplAPIEndpoint, bytes.NewBuffer(data))
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

	return io.ReadAll(resp.Body)
}

func (d *DeeplClient) parseResponse(body []byte) (string, error) {
	var deeplResp translator.TranslationResponse
	if err := json.Unmarshal(body, &deeplResp); err != nil {
		return "", fmt.Errorf("could not unmarshal response data: %v", err)
	}

	if len(deeplResp.Translations) == 0 {
		return "", fmt.Errorf("no translations returned by DeepL")
	}

	return deeplResp.Translations[0].Text, nil
}

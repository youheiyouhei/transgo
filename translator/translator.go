package translator

type Translator interface {
	Translate(request TranslationRequest) (string, error)
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

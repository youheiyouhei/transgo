package interfaces

type Language interface {
	GetSupportedLanguages() (SupportedLanguages, error)
}

type SupportedLanguage struct {
	Code string
	Name string
}

type SupportedLanguages []SupportedLanguage

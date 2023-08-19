package translator

type Translator interface {
	Translate(texts []string, source string, target string) (string, error)
}

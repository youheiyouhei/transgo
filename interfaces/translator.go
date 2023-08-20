package translator

type Translator interface {
	Translate(texts []string, source, target string) (string, error)
}

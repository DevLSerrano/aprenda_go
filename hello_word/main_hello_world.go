package main

const (
	englishHelloPrefix = "Hello, "
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
	french             = "French"
	english            = "English"
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	return gettingPrefix(language) + name
}

func gettingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
}

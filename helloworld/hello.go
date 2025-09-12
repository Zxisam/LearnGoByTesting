package main

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	germanHelloPrefix  = "Hallo, "
	italianHelloPrefix = "Ciao, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "German":
		prefix = germanHelloPrefix
	case "Italian":
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

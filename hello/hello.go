package main

const (
	englishPrefixWord = "Hello, "
	spanishPrefixWord = "Hola, "
	frenchPrefixWord  = "Bonjour, "

	english = "English"
	spanish = "Spanish"
	french  = "French"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishPrefixWord
	case french:
		prefix = frenchPrefixWord
	default:
		prefix = englishPrefixWord
	}
	return
}

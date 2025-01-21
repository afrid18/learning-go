package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const hindi = "Hindi"
const hindiHelloPrefix = "Namasthe, "

func main() {
	fmt.Println(Hello("Afrid", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	prefix = englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix

	}
	return
}

package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
    spanishPrefix = "Hola, "
    frenchPrefix = "Bonjour, "
    spanish = "Spanish"
    french = "French"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World!"
	}

	return languagePrefix(lang) + name
}

func languagePrefix(lang string) (prefix string) {

    switch lang {
    case spanish:
        prefix = spanishPrefix
    case french:
        prefix = frenchPrefix
    default:
        prefix = englishHelloPrefix
    }

    return
}

func main() {
	fmt.Println(Hello("Antonio", ""))
}

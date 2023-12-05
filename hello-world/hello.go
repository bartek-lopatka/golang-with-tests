package main

import "fmt"

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		return "Hello, world!"
	}

	return greetingPrefix(language) + name + "!"

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	default:
		prefix = englishPrefix
	case "fr":
		prefix = frenchPrefix
	case "es":
		prefix = spanishPrefix
	}

	return

}
func main() {
	fmt.Println(Hello("Chris", ""))
}

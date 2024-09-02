package main

import (
	"fmt"
)

const (
	spanish        = "Spanish"
	french         = "French"
	engHelloPrefix = "Hello, "
	spnHelloPrefix = "Hola, "
	frnHelloPrefix = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spnHelloPrefix
	case french:
		prefix = frnHelloPrefix
	default:
		prefix = engHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Nihal", ""))
}

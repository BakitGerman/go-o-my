package main

import "fmt"

const (
	rus = "Russian"
	eng = "English"
)

const (
	englishHelloPrefix = "Hello, "
	russianHelloPrefix = "Привет, "
)

const (
	englishHelloName = "World"
	russianHelloName = "Мир"
)

func Hello(name, language string) string {
	if name == "" {
		name = greetingDefaultName(language)
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case rus:
		prefix = russianHelloPrefix
	case eng:
		prefix = englishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func greetingDefaultName(language string) (name string) {
	switch language {
	case rus:
		name = russianHelloName
	case eng:
		name = englishHelloName
	default:
		name = englishHelloName
	}
	return
}

func main() {
	fmt.Println(Hello("Bakit", ""))
}

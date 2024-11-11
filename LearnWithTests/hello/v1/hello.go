package main

import "fmt"

const englishHelloPrefix = "Hello, "
const enlishHelloWorld = "Hello, World"

func hello(name string) string {
	if name == "" {
		return enlishHelloWorld
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(hello("Bakit"))
}

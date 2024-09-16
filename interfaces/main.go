package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}

type spanishBot struct{}

type Bot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	// Very custom logic for generating a english greeting
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(b englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(b spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

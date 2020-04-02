package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type teluguBot struct{}

func main() {
	var eb englishBot
	var tb teluguBot

	printGreeting(eb)
	printGreeting(tb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello There!"
}

func (teluguBot) getGreeting() string {
	return "Namasthe _/\\_"
}

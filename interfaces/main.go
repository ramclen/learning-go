package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb:= englishBot{}
	es:= spanishBot{}
	printGreetings(eb)
	printGreetings(es)
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (englishBot) getGreetings() string{
	return "hi there!"
}

func (spanishBot) getGreetings() string{
	return "Hola!"
}
package main

type englishBot struct {}
type spanishBot struct {}
type bot interface {
	GetGreeting() string
}

func main()  {
	englishBot := englishBot{}
	spanishBot := spanishBot{}

	printGreeting(englishBot)
	printGreeting(spanishBot)
}

func (e englishBot) GetGreeting() string {
	return "Hello, World!"
}

func (s spanishBot) GetGreeting() string {
	return "Hola, Mundo!"
}

func printGreeting(b bot) {
	println(b.GetGreeting())
}


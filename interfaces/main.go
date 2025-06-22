package main

type englishBot struct {}
type spanishBot struct {}
type bot interface {
	GetGreeting() string
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

func main()  {
	englishBot := englishBot{}
	spanishBot := spanishBot{}

	printGreeting(englishBot)
	printGreeting(spanishBot) // This will cause a compile-time error since printGreeting expects an EnglishBot	
}


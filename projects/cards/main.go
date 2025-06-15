package main

func main()  {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	
	// println(cards.toString())
	// hand.print()
	// remainingCards.print()
	cards.saveToFile("my_cards.txt")
}
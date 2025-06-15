package main

func main()  {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	
	// println(cards.toString())
	// hand.print()
	// remainingCards.print()
	cards := newDeckFromFile("my_cards.txt")
	cards.print()	
}
package main

func main()  {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 5)
	
	// println(cards.toString())
	// hand.print()
	// remainingCards.print()
	// cards := newDeckFromFile("my_cards.txt")
	// cards.print()

}
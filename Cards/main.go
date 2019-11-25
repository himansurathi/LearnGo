package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	fmt.Println(card)

	testCard := "Ten of Diamonds"
	fmt.Println(testCard)

	testCard = "Five of Clubs"
	fmt.Println(testCard)

	testCard = newCard()
	fmt.Println(testCard)

	multipleCards := deck{"Five of Clubs", newCard()}
	fmt.Println(multipleCards)

	newMultipleCards := append(multipleCards, "Six of Spades")
	newMultipleCards.print()

	realDeck := newDeck()
	hand, remainingDeck := realDeck.deal(5)
	fmt.Println("***************Hand Drawn*******************")
	hand.print()
	fmt.Println("***************Remaining Deck******************")
	remainingDeck.print()
}

func newCard() string {
	return "Seven of Diamonds"
}

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

	multipleCards := []string{"Five of Clubs", newCard()}
	fmt.Println(multipleCards)

	newMultipleCards := append(multipleCards, "Six of Spades")
	for i, c := range newMultipleCards {
		fmt.Println(i, c)
	}

}

func newCard() string {
	return "Seven of Diamonds"
}

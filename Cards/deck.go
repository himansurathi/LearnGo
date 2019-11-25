package main

import "fmt"

// Create a new type of deck
// which is a slice of string
type deck []string

// Receiver functions will have abbreviated letter for deck
func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{
		"Spades",
		"Diamonds",
		"Clubs",
		"Hearts"}

	cardValues := []string{
		"Ace", "Two", "Three",
		"Four", "Five", "Six",
		"Seven", "Eight", "Nine",
		"Ten", "Jack", "Queen",
		"King"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

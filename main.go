package main

import "fmt"

func main() {
	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "Queen of Diamond")

	for i, card := range cards {

		fmt.Println(i, card)
	}
}

func newCard() string {
	return "five of Diamonds"
}

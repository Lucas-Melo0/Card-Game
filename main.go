package main

import "fmt"

func main() {

	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	fmt.Println([]byte(hand.toString()))
	fmt.Println([]byte(remainingDeck.toString()))

	/* fmt.Println([]byte()) */
}

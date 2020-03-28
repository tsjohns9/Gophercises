package main

import (
	"fmt"

	"github.com/tsjohns9/gophercises/blackjack"
)

// AI implements the AI interface
type AI struct{}

func main() {
	var ai AI
	// setup ai if you need to...

	opts := blackjack.Options{
		Hands: 100,
		Decks: 3,
	}
	game := blackjack.New(opts)
	winnings := game.Play(ai)
	fmt.Println("Our AI won/lost:", winnings)
}

package blackjack

import (
	"fmt"

	"github.com/tsjohns9/gophercises/deck"
)

// type dealer implements the Player interface

type dealerAI struct{}

func (d dealerAI) Bet() int {
	// Dealer does not bet
	return 0
}

func (d dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	score := Score(hand...)
	fmt.Printf("The dealer has %d points.\n", score)
	if score <= 16 {
		return Hit
	}
	return Stand
}

func (d dealerAI) String(hand []deck.Card) string {
	return hand[0].String() + " --> Hidden <--"
}

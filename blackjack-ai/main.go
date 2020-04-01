package main

import (
	"strings"

	"github.com/tsjohns9/gophercises/blackjack/blackjack"
	"github.com/tsjohns9/gophercises/deck"
)

// AI implements the Player interface
type ai struct{}

func (a ai) Bet() int {
	return 1
}

func (a ai) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	pScore := blackjack.Score(hand...)
	if pScore < 17 {
		return blackjack.Hit
	} else {
		return blackjack.Stand
	}
}

func (a ai) String(hand []deck.Card) string {
	strs := make([]string, len(hand))
	for i, c := range hand {
		strs[i] = c.String()
	}
	return strings.Join(strs, ", ")
}

func main() {
	var player ai
	opts := blackjack.Options{
		Cash:   500,
		Decks:  3,
		Hands:  10,
		MinBet: 10,
	}

	game := blackjack.NewGame(opts)
	game.Play(player)
}

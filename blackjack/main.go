package main

import (
	"fmt"
	"strings"

	"github.com/tsjohns9/gophercises/blackjack/blackjack"
	"github.com/tsjohns9/gophercises/deck"
)

type humanPlayer struct{}

func (h humanPlayer) Bet() int {
	return 1
}

func (h humanPlayer) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	for {
		fmt.Printf("Your hand: %v\n\n", h.String(hand))
		fmt.Printf("The dealers hand: %v\n\n", dealer.String())
		fmt.Printf("You have %d points. The dealer has %d points.\n", blackjack.Score(hand...), blackjack.Score(dealer))
		fmt.Println("Your turn. Will you hit (h), or stand (s) ?")

		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return blackjack.Hit
		case "s":
			return blackjack.Stand
		default:
			break
		}
	}
}

func (h humanPlayer) String(hand []deck.Card) string {
	strs := make([]string, len(hand))
	for i, c := range hand {
		strs[i] = c.String()
	}
	return strings.Join(strs, ", ")
}

func main() {
	var player humanPlayer
	opts := blackjack.Options{
		Cash:   500,
		Decks:  3,
		Hands:  10,
		MinBet: 10,
	}

	game := blackjack.NewGame(opts)
	game.Play(player)
}

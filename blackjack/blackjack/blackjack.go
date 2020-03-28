package blackjack

import (
	"fmt"

	"github.com/tsjohns9/gophercises/deck"
)

type turn uint8

const (
	_ turn = iota
	playerTurn
	dealerTurn
)

// Options are options to start the game
type Options struct {
	Cash   int
	Decks  int
	Hands  int
	MinBet int
}

// Player represents methods for a blackjack player, human or AI.
type Player interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Move
	String([]deck.Card) string
	// Results(hand [][]deck.Card, dealer []deck.Card)
}

// Move is a move
type Move func(*Game)

// Game is a game
type Game struct {
	balance    int
	dealerAI   dealerAI
	dealerHand []deck.Card
	deck       []deck.Card
	decks      int
	playerBet  int
	playerHand []deck.Card
	turn       turn
}

// Play starts the game with the given player
func (g Game) Play(player Player) {
	for g.turn == playerTurn {
		move := player.Play(g.playerHand, g.dealerHand[0])
		if move == nil {
			return
		}
		move(&g)
	}
	for g.turn == dealerTurn {
		move := g.dealerAI.Play(g.dealerHand, g.dealerHand[0])
		move(&g)
	}
	endGame(&g)
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

// NewGame starts the game based on the options
func NewGame(options Options) Game {
	cards := deck.NewDeck(deck.MultiDeck(options.Decks), deck.Shuffle)

	var game Game
	var card deck.Card
	var playerHand, dealerHand []deck.Card

	for i := 0; i < 2; i++ {
		for _, hand := range []*[]deck.Card{&playerHand, &dealerHand} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	game.balance = options.Cash
	game.dealerAI = dealerAI{}
	game.dealerHand = dealerHand
	game.deck = cards
	game.decks = options.Decks
	game.playerBet = options.MinBet
	game.playerHand = playerHand
	game.turn = playerTurn

	return game
}

// Hit draws another card
func Hit(g *Game) {
	card, deck := draw(g.deck)
	g.deck = deck
	if g.turn == playerTurn {
		g.playerHand = append(g.playerHand, card)
	} else if g.turn == dealerTurn {
		g.dealerHand = append(g.dealerHand, card)
	}
}

// Stand ends the round for the player
func Stand(g *Game) {
	if g.turn == playerTurn {
		g.turn = dealerTurn
		fmt.Println("*** DEALERS TURN ***")
	} else {
		g.turn = playerTurn
	}
}

// Score counts the score
func Score(hand ...deck.Card) int {
	sum := 0
	for _, h := range hand {
		sum += min(int(h.Rank))
	}
	return sum
}

func min(n int) int {
	if n < 10 {
		return n
	}
	return 10
}

func endGame(g *Game) {
	pScore, dScore := Score(g.playerHand...), Score(g.dealerHand...)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
		g.balance--
	case dScore > 21:
		fmt.Println("Dealer busted")
		g.balance++
	case pScore > dScore:
		fmt.Printf("You win! You have %d points. The dealer has %d points.", pScore, dScore)
		g.balance++
	case dScore > pScore:
		fmt.Printf("You lose. You have %d points. The dealer has %d points.", pScore, dScore)
		g.balance--
	case dScore == pScore:
		fmt.Println("Draw")
	}
}

package blackjack

import (
	"errors"
	"fmt"

	"github.com/tsjohns9/gophercises/deck"
)

type turn uint8

var (
	errBust = errors.New("You busted")
)

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
type Move func(*Game) error

// Game is a game
type Game struct {
	balance    int
	dealerAI   dealerAI
	dealerHand []deck.Card
	deck       []deck.Card
	decks      int
	hands      int
	playerBet  int
	playerHand []deck.Card
	turn       turn
	wins       int
}

// Play starts the game with the given player
func (g Game) Play(player Player) {
	for i := 0; i < g.hands; i++ {
		fmt.Printf("Round %d:\n", i+1)
		for g.turn == playerTurn {
			move := player.Play(g.playerHand, g.dealerHand[0])
			if move == nil {
				return
			}
			err := move(&g)
			if pScore := Score(g.playerHand...); pScore == 21 {
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~ You have 21 points.")
				break
			}
			if err != nil {
				break
			}
		}
		for g.turn == dealerTurn {
			move := g.dealerAI.Play(g.dealerHand, g.dealerHand[0])
			move(&g)
		}
		endGame(&g)
		g.deck = shuffle(g.decks)
		playerHand, dealerHand := deal(g.deck)
		g.playerHand = playerHand
		g.dealerHand = dealerHand
	}
	fmt.Printf("You won %d out of %d hands.\n", g.wins, g.hands)
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

// NewGame starts the game based on the options
func NewGame(options Options) Game {
	cards := shuffle(options.Decks)

	var game Game
	playerHand, dealerHand := deal(cards)

	game.balance = options.Cash
	game.dealerAI = dealerAI{}
	game.dealerHand = dealerHand
	game.deck = cards
	game.decks = options.Decks
	game.hands = options.Hands
	game.playerBet = options.MinBet
	game.playerHand = playerHand
	game.turn = playerTurn

	return game
}

func deal(cards []deck.Card) ([]deck.Card, []deck.Card) {
	var card deck.Card
	var playerHand, dealerHand []deck.Card
	for i := 0; i < 2; i++ {
		for _, hand := range []*[]deck.Card{&playerHand, &dealerHand} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	return playerHand, dealerHand
}

func shuffle(decks int) []deck.Card {
	return deck.NewDeck(deck.MultiDeck(decks), deck.Shuffle)
}

// Hit draws another card
func Hit(g *Game) error {
	card, deck := draw(g.deck)
	g.deck = deck
	if g.turn == playerTurn {
		g.playerHand = append(g.playerHand, card)
	} else if g.turn == dealerTurn {
		g.dealerHand = append(g.dealerHand, card)
	}
	if Score(g.playerHand...) > 21 {
		return errBust
	}
	return nil
}

// Stand ends the round for the player
func Stand(g *Game) error {
	if g.turn == playerTurn {
		g.turn = dealerTurn
		fmt.Println("*** DEALERS TURN ***")
	} else {
		g.turn = playerTurn
	}
	return nil
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
		fmt.Printf("--------You busted with %d points.--------\n\n", pScore)
		g.balance--
	case dScore > 21:
		fmt.Printf("--------Dealer busted. with %d points. You Won with %d points!--------\n\n", dScore, pScore)
		g.balance++
		g.wins++
	case pScore > dScore:
		fmt.Printf("--------You win! You have %d points. The dealer has %d points.--------\n\n", pScore, dScore)
		g.balance++
		g.wins++
	case dScore > pScore:
		fmt.Printf("--------You lose. You have %d points. The dealer has %d points.--------\n\n", pScore, dScore)
		g.balance--
	case dScore == pScore:
		fmt.Println("--------Draw--------")
	}
}

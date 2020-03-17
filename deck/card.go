//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit is a suit
type Suit uint8

// Rank is a rank
type Rank uint8

// Suits
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

// Ranks
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Card is a card
type Card struct {
	Suit
	Rank
}

var suits = [...]Suit{Spade, Diamond, Club, Heart}

func absRank(c Card) int {
	return int(c.Suit)*int(King) + int(c.Rank)
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// NewDeck creates a new deck
func NewDeck(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := Ace; rank <= King; rank++ {
			card := Card{Suit: suit, Rank: rank}
			cards = append(cards, card)
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

// DefaultSort sorts
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, LeastToGreatest(cards))
	return cards
}

// LeastToGreatest sorts cards from least to greatest
func LeastToGreatest(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

// CustomSort is a custom sort
func CustomSort(less func(card []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Shuffle shuffles
func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// perm := r.Perm(len(cards))
	for i := range cards {
		newPosition := r.Intn(len(cards) - 1)
		cards[i], cards[newPosition] = cards[newPosition], cards[i]
	}
	return cards
}

// AddJoker adds n jokers
func AddJoker(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker})
		}
		return cards
	}
}

// Filter filters the deck of cards based on the provided func
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var newDeck []Card
		for _, c := range cards {
			if !f(c) {
				newDeck = append(newDeck, c)
			}
		}
		return newDeck
	}
}

// MultiDeck returns n decks
func MultiDeck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card

		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}

		return ret
	}
}

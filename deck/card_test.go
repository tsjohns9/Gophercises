package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Printf("%v\n", Card{Rank: Ace, Suit: Diamond})

	// Output:
	// Ace of Diamonds
}

func TestNew(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 52 {
		t.Error("Wrong number of cards. Received", len(cards))
	}
}

func TestDefaultSort(t *testing.T) {
	cards := NewDeck(DefaultSort)
	if cards[0].String() != "Ace of Spades" {
		t.Error("Error: Received", cards[0].String(), "Expected: Ace of Spades")
	}
}

func TestAddJoker(t *testing.T) {
	jokerTotal := 2
	cards := NewDeck(AddJoker(jokerTotal))
	if len(cards) != 54 {
		t.Errorf("Error: expected a total card count of %d but received %d\n", 54, len(cards))
	}

	count := 0
	for _, c := range cards {
		if c.String() == "Joker" {
			count++
		}
	}
	if count != jokerTotal {
		t.Errorf("Error: Expected %d Jokers but received %d\n", jokerTotal, count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(c Card) bool {
		return c.Rank == Two
	}

	cards := NewDeck(Filter(filter))

	if len(cards) != 48 {
		t.Errorf("Error: Expected 48 cards but received %d\n", len(cards))
	}

	for _, c := range cards {
		if c.Rank == Two {
			t.Error("Error: Expected Cards with a rank of Two to be filtered.")
			break
		}
	}
}

func TestDeck(t *testing.T) {
	cards := NewDeck(MultiDeck(3))
	if len(cards) != 52*3 {
		t.Errorf("Error: Expected %d cards but received %d", 52*3, len(cards))
	}
}

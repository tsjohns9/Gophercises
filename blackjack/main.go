package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tsjohns9/gophercises/deck"
)

// Hand is a blackjack hand
type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i, c := range h {
		strs[i] = c.String()
	}
	return strings.Join(strs, ", ")
}

// DealerString prints the dealers hand
func (h Hand) DealerString() string {
	return h[0].String() + " --> Hidden <--"
}

// Score gets the score of the hand
func (h Hand) Score() int {
	sum := 0
	for _, card := range h {
		i := int(card.Rank)
		if i < 10 {
			sum += int(card.Rank)
		} else {
			sum += 10
		}
	}
	return sum
}

func main() {
	cards := deck.NewDeck(deck.MultiDeck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Your hand: %v\n\n", player)
	fmt.Println("Your turn. Will you hit (h), or stand (s) ?")

	score := player.Score()
	dealerScore := dealer.Score()

	if dealerScore == 21 {
		fmt.Println("The dealer hit 21 points and you lost!")
		return
	}

	var nextLine string
	for nextLine != "s" && score < 21 {
		fmt.Printf("You have %d points. The dealer has %d points.\n", score, dealerScore)

		bts, _, _ := reader.ReadLine()

		nextLine = string(bts)
		if nextLine == "h" {
			switch nextLine {
			case "h":
				card, cards = draw(cards)
				player = append(player, card)
				score = player.Score()
				fmt.Printf("\nYour next card is %v. You now have %d points.\n\n", card, player.Score())
			}
		}
	}
	if nextLine == "s" && score < 21 && dealerScore < 17 {
		for dealerScore < 17 {
			card, cards = draw(cards)
			dealer = append(dealer, card)
			dealerScore = dealer.Score()
			fmt.Printf("The dealer drew the %v and now has %d points.\n\n", card, dealerScore)
			if dealerScore < 17 {
				time.Sleep(2 * time.Second)
			}
		}
	}
	if dealerScore > 21 {
		fmt.Printf("The dealer busted with %d points.\n", dealerScore)
		return
	}
	if score == 21 {
		fmt.Println("You won with 21 points!")
		return
	}
	if score > 21 {
		fmt.Printf("You busted with %d points.\n", score)
		return
	}

	if score < 21 && score > dealerScore {
		fmt.Printf("You beat the dealer with %d points. The dealer had %d points.\n", score, dealerScore)
	} else {
		fmt.Printf("The dealer won with %d points. You have %d points\n", dealerScore, score)
	}

}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

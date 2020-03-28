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

// func main() {
// 	cards := deck.NewDeck(deck.MultiDeck(3), deck.Shuffle)
// 	var card deck.Card
// 	var player, dealer blackjack.Hand

// 	for i := 0; i < 2; i++ {
// 		for _, hand := range []*blackjack.Hand{&player, &dealer} {
// 			card, cards = blackjack.Draw(cards)
// 			*hand = append(*hand, card)
// 		}
// 	}
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Printf("Your hand: %v\n\n", player)
// 	fmt.Println("Your turn. Will you hit (h), or stand (s) ?")

// 	score := player.Score()
// 	dealerScore := dealer.Score()

// 	if dealerScore == 21 {
// 		fmt.Println("The dealer hit 21 points and you lost!")
// 		return
// 	}

// 	var nextLine string
// 	for nextLine != "s" && score < 21 {
// 		fmt.Printf("You have %d points. The dealer has %d points.\n", score, dealerScore)

// 		bts, _, _ := reader.ReadLine()

// 		nextLine = string(bts)
// 		if nextLine == "h" {
// 			switch nextLine {
// 			case "h":
// 				card, cards = blackjack.Draw(cards)
// 				player = append(player, card)
// 				score = player.Score()
// 				fmt.Printf("\nYour next card is %v. You now have %d points.\n\n", card, player.Score())
// 			}
// 		}
// 	}
// 	if nextLine == "s" && score < 21 && dealerScore < 17 {
// 		for dealerScore < 17 {
// 			card, cards = blackjack.Draw(cards)
// 			dealer = append(dealer, card)
// 			dealerScore = dealer.Score()
// 			fmt.Printf("The dealer drew the %v and now has %d points.\n\n", card, dealerScore)
// 			if dealerScore < 17 {
// 				time.Sleep(2 * time.Second)
// 			}
// 		}
// 	}
// 	if dealerScore > 21 {
// 		fmt.Printf("The dealer busted with %d points.\n", dealerScore)
// 		return
// 	}
// 	if score == 21 {
// 		fmt.Println("You won with 21 points!")
// 		return
// 	}
// 	if score > 21 {
// 		fmt.Printf("You busted with %d points.\n", score)
// 		return
// 	}

// 	if score < 21 && score > dealerScore {
// 		fmt.Printf("You beat the dealer with %d points. The dealer had %d points.\n", score, dealerScore)
// 	} else {
// 		fmt.Printf("The dealer won with %d points. You have %d points\n", dealerScore, score)
// 	}
// }

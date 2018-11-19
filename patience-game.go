package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	suit  string
	value string
}

func (card Card) String() string {
	return card.value + " of " + card.suit
}

var suits [4]string = [4]string{"Spades", "Clubs", "Diamonds", "Hearts"}
var values [13]string = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	deck := make([]Card, 0, 52)
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{suit, value})
		}
	}
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	fmt.Println("First card is: ", deck[0])
}

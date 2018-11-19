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
	
	deck := makeDeck()
	shuffled := shuffledDeck(deck)
	
	fmt.Println("First card is: ", shuffled[0])
}

func makeDeck() []Card {
	deck := make([]Card, 0, 52)
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{suit, value})
		}
	}
	return deck
}

func shuffledDeck(originalDeck []Card) []Card {
	copiedDeck := make([]Card, len(originalDeck))
	copy(copiedDeck, originalDeck)
	rand.Shuffle(len(copiedDeck), func(i, j int) {
		copiedDeck[i], copiedDeck[j] = copiedDeck[j], copiedDeck[i]
	})
	return copiedDeck
}

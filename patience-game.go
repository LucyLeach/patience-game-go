package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value string
}

func (card Card) String() string {
	return card.Value + " of " + card.Suit
}

var suits [4]string = [4]string{"Spades", "Clubs", "Diamonds", "Hearts"}
var values [13]string = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	numSimulations := 400000

	deck := makeDeck()

	winCount := 0
	for i := 0; i < numSimulations; i++ {
		shuffled := shuffledDeck(deck)
		win := playPatience(shuffled)
		if win {
			winCount++
		}
	}

	winPercent := 100.0 * float64(winCount) / float64(numSimulations)
	fmt.Printf("Won %v of %v games, %v%%", winCount, numSimulations, winPercent)
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

func playPatience(deck []Card) bool {
	piles := make(map[string][]Card)
	for i, value := range values {
		piles[value] = deck[i*4: i*4 + 4]
	}
	
	currentCard := removeBottomCard(piles, "King")
	for len(piles[currentCard.Value]) > 0 {
		currentCard = removeBottomCard(piles, currentCard.Value)
	}
	
	remainingCards := 0
	for _, pile := range piles {
		remainingCards += len(pile)
	}
	
	return remainingCards == 0
}

func removeBottomCard(piles map[string][]Card, value string) Card {
	currentPile := piles[value]
	lastCard := currentPile[len(currentPile) - 1]
	piles[value] = currentPile[:len(currentPile) - 1]
	return lastCard
}

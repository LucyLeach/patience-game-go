package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
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
	numSimulations := 10000000

	numProcs := runtime.GOMAXPROCS(0)
	//May result in a couple more or less than requested
	numSimulationsPerRoutine := int(math.Round(float64(numSimulations) / float64(numProcs)))

	deck := makeDeck()

	c := make(chan int)
	for i := 0; i < numProcs; i++ {
		go playNGamesOfPatience(deck, numSimulationsPerRoutine, c)
	}

	winCount := 0
	for i := 0; i < numProcs; i++ {
		winCount += <-c
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

func playNGamesOfPatience(deck []Card, n int, c chan int) {
	winCount := 0
	for i := 0; i < n; i++ {
		if playPatience(deck) {
			winCount++
		}
	}
	c <- winCount
}

func playPatience(orderedDeck []Card) bool {
	deck := shuffledDeck(orderedDeck)

	piles := make(map[string][]Card)
	for i, value := range values {
		piles[value] = deck[i*4 : i*4+4]
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
	lastCard := currentPile[len(currentPile)-1]
	piles[value] = currentPile[:len(currentPile)-1]
	return lastCard
}

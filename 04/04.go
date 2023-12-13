package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Scratchcard struct {
	count          int
	winningNumbers []int
	pickedNumbers  []int
}

func (s Scratchcard) Matches() []int {
	var matches []int

	for _, n := range s.pickedNumbers {
		// Number is a winning number
		if slices.Contains(s.winningNumbers, n) {
			matches = append(matches, n)
		}
	}

	return matches
}

func (s Scratchcard) Worth() int {
	var worth int

	if len(s.Matches()) != 0 {
		worth = 1
		// Calculate card worth and skip first
		for range s.Matches()[1:] {
			worth = worth * 2
		}
	} else {
		worth = 0
	}

	return worth
}

type Cardpile []Scratchcard

func main() {
	solve(input)
}

func solve(input string) {
	var pileWorth int
	var cardPile Cardpile
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	// Parse winning numbers and picked numbers out of input
	for _, card := range lines {
		var winningNumbers []int
		var myNumbers []int

		l := strings.Split(card, ":")
		numbers := strings.Split(l[1], "|")

		for _, numberString := range strings.Fields(numbers[0]) {
			numberInt, _ := strconv.Atoi(numberString)
			winningNumbers = append(winningNumbers, numberInt)
		}

		for _, numberString := range strings.Fields(numbers[1]) {
			numberInt, _ := strconv.Atoi(numberString)
			myNumbers = append(myNumbers, numberInt)
		}

		card := Scratchcard{count: 1, winningNumbers: winningNumbers, pickedNumbers: myNumbers}
		cardPile = append(cardPile, card)

		pileWorth += card.Worth()
	}

	fmt.Printf("Part 1: %v\n", pileWorth)

	// Go through the card pile and track copies
	for id, card := range cardPile {
		matchingCount := len(card.Matches())

		for i := id + 1; i <= (id + matchingCount); i++ {
			cardPile[i].count += card.count
		}
	}

	// Calculate total amount of cards with copies
	total := 0
	for _, card := range cardPile {
		total = total + card.count
	}

	fmt.Printf("Part 2: %v\n", total)
}

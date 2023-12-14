package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HighCard int = iota // starts from 0 continues +1 for each subsequent const
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards string
	Bids  int
	Type  int
}

func main() {
	file, err := os.Open("../test_input")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := make([]Hand, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{Cards: parts[0], Bids: bid, Type: classifyHand(parts[0])})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading from file: %s", err)
	}

	// order by type
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].Type != hands[j].Type {
			return hands[i].Type > hands[j].Type
		}
		// order by first biggest card if same type
		return compareCards(hands[i].Cards, hands[j].Cards)
	})

	totalWinnings := 0
	for i, hand := range hands {
		rank := len(hands) - i
		totalWinnings += hand.Bids * rank
	}

	fmt.Println("Total winnings:", totalWinnings)
}

func classifyHand(hand string) int {
	counts := make(map[rune]int)
	for _, card := range hand {
		counts[card]++ // key is char/rune, val is int: how many times it is present!
	}
	fmt.Println("map[rune]int: ", counts)

	switch len(counts) {
	case 1:
		return FiveOfAKind
	case 2:
		for _, count := range counts {
			if count == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	case 3:
		for _, count := range counts {
			if count == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPairs
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func compareCards(hand1, hand2 string) bool {
	cardOrder := map[rune]int{'2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'T': 9, 'J': 10, 'Q': 11, 'K': 12, 'A': 13}
	for i := 0; i < len(hand1); i++ {
		if cardOrder[rune(hand1[i])] != cardOrder[rune(hand2[i])] {
			return cardOrder[rune(hand1[i])] > cardOrder[rune(hand2[i])]
		}
	}
	return false
}

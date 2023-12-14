package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Card represents a card with its value and a flag indicating if it's a Joker.
type Card struct {
	Value string
	Joker bool
}

// Hand represents a poker hand.
type Hand struct {
	Cards   []Card
	Bid     int
	Ranking int
}

// parseHand parses a string representation of a hand and its bid.
func parseHand(input string) Hand {
	parts := strings.Fields(input)
	cards := make([]Card, len(parts[0]))
	for i, val := range parts[0] {
		cards[i] = Card{Value: string(val), Joker: string(val) == "J"}
	}
	bid, _ := strconv.Atoi(parts[1])
	return Hand{Cards: cards, Bid: bid}
}

// evaluateHand evaluates the hand and sets its ranking based on the game rules.
func evaluateHand(hand *Hand) {
    cardOrder := "AKQJT98765432J"
    counts := make(map[rune]int)
    jokers := 0

    // Count the occurrences of each card and identify Jokers
    for _, card := range hand.Cards {
        if card.Joker {
            jokers++
        } else {
            counts[rune(card.Value[0])]++
        }
    }

    // Determine the best hand possible
    types := make([]int, len(cardOrder))
    for i, card := range cardOrder {
        types[i] = counts[rune(card)]
    }

    // Adjust for Jokers and determine hand type
    handType := determineHandType(types, jokers)

    // Set the ranking based on the hand type
    hand.Ranking = handType
}

// determineHandType determines the type of hand, adjusted for Jokers.
// This function should return an int representing the hand's rank.
func determineHandType(types []int, jokers int) int {
    // Implement the logic to determine the hand type considering the rules
    // This function should analyze the 'types' slice and 'jokers' count
    // and return an integer representing the hand's rank

    // Placeholder for hand type - Needs actual implementation based on Camel Cards rules
    return 1 // Placeholder rank
}

// Additional logic for tie-breaking and hand evaluation should be added as needed


// ByRank implements sort.Interface for []Hand based on the Ranking field.
type ByRank []Hand

func (a ByRank) Len() int           { return len(a) }
func (a ByRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool { return a[i].Ranking < a[j].Ranking }

// calculateTotalWinnings calculates the total winnings based on hand ranks and bids.
func calculateTotalWinnings(hands []Hand) int {
	sort.Sort(ByRank(hands))
	total := 0
	for rank, hand := range hands {
		total += (rank + 1) * hand.Bid
	}
	return total
}

func main() {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	hands := []Hand{}
	for _, line := range strings.Split(input, "\n") {
		hand := parseHand(line)
		evaluateHand(&hand)
		hands = append(hands, hand)
	}

	totalWinnings := calculateTotalWinnings(hands)
	fmt.Println("Total winnings:", totalWinnings)
}

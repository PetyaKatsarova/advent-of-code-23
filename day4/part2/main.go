package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("../puzzle_input")
	file, err := os.Open("../test_input")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cards := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, line)
	}

	totalCards := processCards(cards)
	fmt.Println("Total scratchcards:", totalCards)
}

func processCards(cards []string) int {
	totalCount := 0
	queue := make([]string, len(cards))
	copy(queue, cards)

	for len(queue) > 0 {
		card := queue[0]
		queue = queue[1:]
		totalCount++

		_, wins := calculateCard(card)
		cardIndex, _ := strconv.Atoi(strings.Fields(strings.Split(card, ":")[0])[1])
		for i := 0; i < wins; i++ {
			if cardIndex + i < len(cards) {
				queue = append(queue, cards[cardIndex + i])
			}
		}
	}

	return totalCount
}

func calculateCard(line string) (int, int) {
	str := strings.Split(line, ":")
	cardNumberStr := strings.Fields(str[0])[1]
	cardNumber, err := strconv.Atoi(cardNumberStr)
	if err != nil {
		log.Fatalf("error in strconv.Atoi(cardNumberStr)")
	}

	numbers := strings.Split(str[1], "|")
	winningNums := strings.Fields(numbers[0])
	gameNums := strings.Fields(numbers[1])

	wins := 0
	for _, num := range winningNums {
		for _, n := range gameNums {
			if num == n {
				wins++
				break
			}
		}
	}
	return cardNumber, wins
}

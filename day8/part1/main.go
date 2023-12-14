package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

)

type Letters struct {
	Result	string
	Left	string
	Right	string
}

func main() {
	file, err := os.Open("../puzzle_input")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner		:= bufio.NewScanner(file)
	scanner.Scan()
	directions	:= scanner.Text()
	lettersMap		:= make(map[string]Letters)
	for scanner.Scan() {
		line := scanner.Text()
		splittedLetters := strings.Split(line, " = ")
		if len(splittedLetters) > 1 {
			str	:= strings.Split(splittedLetters[1], ", ")
			left	:= str[0][1:]
			right	:= str[1][:len(str[1]) - 1]
			result	:= splittedLetters[0]
			letter	:= Letters{Result: result, Left: left, Right: right}
			lettersMap[result] = letter
		}
	}
	fmt.Println("the result is: ", followTheRoute(lettersMap, directions))
}

func followTheRoute(letters map[string]Letters, directions string) int {
	currLetters := "AAA"
	if _, ok := letters[currLetters]; !ok {
		log.Fatalf("Starting combination %s not found", currLetters)
	}
	currCombination := letters[currLetters]
	moves := 0

	for currLetters != "ZZZ" {
		for _, val := range directions {
			if val == 'L' {
				currLetters = currCombination.Left 
			} else if val == 'R' {
				currLetters = currCombination.Right
			}

			var ok bool
			currCombination, ok = letters[currLetters]
			if !ok {
				log.Fatalf("Combination %s not found", currLetters)
			}
			moves++
		}
	}
	return moves
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../puzzle_input")
	if err != nil {	log.Fatalf("failed to open file: %s", err)}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += calculateCard(line)
	}
	fmt.Println(sum)
}

func calculateCard(line string) int {
	str := strings.Split(line, ":")
	numbers := strings.Split(str[1], "|")
	winningNums := strings.Fields(numbers[0])
	gameNums    := strings.Fields(numbers[1])
		/*  in order to get the nums and not the individual runes: 
    splits the string s around each instance of one or more consecutive white space characters,
	 as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains
	  only white space
	*/

	points  := 0
	for _, num := range winningNums {
		for _, n := range gameNums {
			if num == n {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
	}
	fmt.Println("--------------------------------------")
	return points
}

/*
When you range over a string in Go, you get individual bytes (runes), not words or numbers separated
 by spaces. Since numbers[0] is a string, for _, num := range numbers[0] iterates over each character
  in the string, including spaces.
To iterate over each number in the string "41 48 83 86 17", you first need to split this string 
into separate numbers using strings.Fields which splits based on whitespace. Here's the corrected
 function:

 -- contains is defined to check if a slice of strings (gameNums) contains a specific string.
*/
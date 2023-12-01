package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println(sumTheNums())
}

func sumTheNums() int64 {
	file, err := os.Open("../puzzle_input") // get pointer to the file object
	// file, err := os.Open("../test_input") // 29, 83, 13, 24, 42, 14, and 76 total is: 281
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	counter := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line %d: %s", counter, line)
		d1, d2 := findDigits(line)
		if d1 != -1 && d2 != -1 {
			total += d1*10 + d2
		}
		fmt.Println("\n", d1, d2)
		counter++
	}
	return int64(total)
}

func findDigits(line string) (int, int) {
    digitsMap := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
    firstDigit, lastDigit := -1, -1
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    currentWord := ""

    for _, char := range line {
        if char >= '0' && char <= '9' {
            digit := int(char - '0')
            if firstDigit == -1 {
                firstDigit = digit
            }
            lastDigit = digit
            currentWord = "" // Clear current word as we found a digit
        } else {
			currentWord += string(char)
			for i, _ := range words {
				if strings.Contains(currentWord, words[i]) {
					currentWord = words[i]
				}
			}
            if val, ok := digitsMap[currentWord]; ok {
                if firstDigit == -1 {
                    firstDigit = val
                }
                lastDigit = val
				currentWord = ""
            }
        }
		
    }

    return firstDigit, lastDigit
}



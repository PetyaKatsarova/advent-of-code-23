package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(sumTheNums())
}

func sumTheNums() int64 {
	file, err := os.Open("../puzzle_input") // get pointer to the file object
	if err != nil { log.Fatalf("failed to open file: %s", err) }
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file) // representation of the scanner object, not the contents of the file. 
	//scanner is an instance of bufio.Scanner, which is a struct in Go, and when you print it directly, it outputs the internal representation of this struct, not the text it reads from the file.
	//The bufio.Scanner type is designed to read input (like files) and break it into lines or words. It's typically used with a loop to read through all lines of a file. However, just printing the scanner variable itself doesn't trigger this reading process. Instead, you would normally use scanner.Scan() in a loop to read each line, and scanner.Text() to get the text of the current line.
	counter := 1
	for scanner.Scan() {
		line := scanner.Text() // like read line in C
		// fmt.Print(counter, ": ")
		d1, d2 := findDigits(line)
		if d1 != -1 && d2 != -1 {
			num, err := strconv.Atoi(fmt.Sprintf("%d%d", d1, d2))
			if err != nil {	log.Fatalf("error converting to int: %s", err) }
			total += num
		}
		// fmt.Println(d1, d2)
		counter++
	}
	return int64(total)
}

func findDigits(s string) (int, int) {
	firstDigit, lastDigit := -1, -1
		for _, char := range s {
			if unicode.IsDigit(char) {
				if firstDigit == -1 {
					firstDigit = int(char - '0')
				}
				lastDigit = int(char - '0') // keeps looping and allocates the last digit
			}
		}
		return firstDigit, lastDigit
}
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
	fmt.Println("total is:", calcTotal())
}

func calcTotal() int {
	file, err := os.Open("../puzzle_input")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameTries := strings.Split(line, ":")[1]
		bla,_ := highestColorNumMultiplied(gameTries)
		total += bla
		
	}
	return total
 }

func highestColorNumMultiplied(line string) (int, error) {
	highestGrene := 1
	highestBlue := 1
	highestRed := 1
	colorNum := strings.Split(strings.ReplaceAll(line, ";", ","), ",")
	
	fmt.Println(colorNum)
	for _, val := range colorNum {
		pair := strings.Fields(strings.TrimSpace(val))
		if len(pair) != 2 {
			return 0, fmt.Errorf("invalid pair: %s", val)
		}
		num, err := strconv.Atoi(pair[0])
		if err != nil {   return 0, fmt.Errorf("error converting string to int: %s", err) }
		switch(pair[1]) {
		case "green":
			highestGrene = getHighestNum(num, highestGrene)
		case "blue":
			highestBlue = getHighestNum(num, highestBlue)
		case "red":
			highestRed = getHighestNum(num, highestRed)
		}
	}
	fmt.Printf("red: %d, blue: %d, green: %d; multiplied result: %d ",  highestBlue, highestGrene, highestRed, highestBlue * highestGrene * highestRed)
	return highestBlue * highestGrene * highestRed, nil
}
//part1> Aho Corasick ook

func getHighestNum(num int, highestNum int) int{
	if num > highestNum {
		highestNum = num
	}
	return highestNum
}

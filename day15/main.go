package main

import (
	"fmt"
	"strings"
	"time"

	"main.go/utils"
)

func calcNumPerStr(inputStr string) int {
	result := 0
	for i := 0; i < len(inputStr); i++ {
		fmt.Println("result:", result)
		result += int(inputStr[i])
		fmt.Println(" +++ result:", result)
		result = result * 17
		fmt.Println("* 17:", result)
		result = result % 256
		fmt.Println("% 256", result)
	} 
	return result
}

func Part1(input []string) int {
	inp := input[0]
	splittedInp := strings.Split(inp, ",")
	result := 0

	for _, val := range splittedInp {
		result += calcNumPerStr(val)
	}

	return result
}

func main() {
	input := utils.ReadFile("puzzle_input")
	for _, val := range input {
		fmt.Println("line; ", val)
	}

	start := time.Now()
	fmt.Println("result part1: ", Part1(input))
	fmt.Println("time: ", time.Since(start))
	fmt.Println("hello day 15")
}
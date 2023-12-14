package main

import (
	"fmt"
	"strings"
	"time"
	"bufio"
	"os"
)

type Network struct {
	Instructions string
	Nodes		 map[string][2]string
}

func parseLine(line string) (head string, children [2]string)  {
	parts := strings.Split(line, " = ")

	head = parts[0]
	childrenTrim := strings.Trim(parts[1], "()")
	childrenParts := strings.Split(childrenTrim, ", ")
	children  = [2]string { childrenParts[0], childrenParts[1]}

	return head, children
}

func parseInput(input []string) Network {
	instructions := input[0]

	nodes := map[string][2]string{}
	for _, line := range input[2:] {
		head, children := parseLine(line)
		nodes[head] = children
	}

	networt := Network {
		Instructions: instructions,
		Nodes:        nodes,
	}
	return networt
}

func greatestCommonDenominator(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func leastCommonMultiple(a, b int) int {
	return (a*b) / greatestCommonDenominator(a, b)
}

func lcmSlice(nums []int) int {
	if len(nums) == 0 {	return 0 }

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = leastCommonMultiple(res, nums[i])
	}
	return res
}

func Part2(input []string) int {
	network := parseInput(input)

	starts := []string{}
	for key := range network.Nodes {
		lastIndex := len(key) - 1
		if key[lastIndex] == 'A' {
			starts = append(starts, key)
		}
	}

	steps := make([]int, len(starts))
	instructionsLen := len(network.Instructions)
	for i := 0; i < len(starts); i++ {
		el := starts[i]
		lastIndex := len(el) - 1
		for el[lastIndex] != 'Z' {
			instruction := network.Instructions[steps[i]%instructionsLen] // clever way to keep looping
			// through the instructions after reached the end
			if instruction == 'L' {
				el = network.Nodes[el][0]
			} else {
				el = network.Nodes[el][1]
			}
			steps[i]++
		}
	}
	return lcmSlice(steps)
}

	func ReadFile(fileName string) []string {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		return lines
	}

	func main() {
		start1 := time.Now()
		fmt.Println("result is: ", Part2(ReadFile("../day8/test_input")))
		fmt.Println(time.Since(start1))
	}
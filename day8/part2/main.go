package main
// learning from: https://github.com/OscarBrugne/AdventOfCode/blob/main/utils/utils.go

import (
	"fmt"
	"strings"
	"time"
	"bufio"
	"os"
)

type Network struct {
	Instructions string
	Nodes        map[string][2]string
}

func parseInput(input []string) Network {
	instructions := input[0]

	nodes := map[string][2]string{}
	for _, line := range input[2:] { // from index 2 because of the free space line
		head, children := parseLine(line)
		nodes[head] = children
	}

	network := Network{
		Instructions: instructions,
		Nodes:        nodes,
	}
	return network
}

func parseLine(line string) (head string, children [2]string) {
	parts := strings.Split(line, " = ")

	head = parts[0]
	childrenTrim := strings.Trim(parts[1], "()")
	childrenParts := strings.Split(childrenTrim, ", ")
	children = [2]string{childrenParts[0], childrenParts[1]}

	return head, children
}

func gcd(a, b int) int { // gratest common denominator
	for b != 0 {
		a, b = b, a%b // Euclidean algorithm for finding the greatest common divisor (GCD) of two numbers. 
	}
	return a
}

func lcm(a, b int) int { // least common multiple (LCM) of two integers a and b. 
	//The LCM is the smallest positive integer that is divisible by both a and b. 
	return (a * b) / gcd(a, b)
}

func lcmSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = lcm(res, nums[i])
	}

	return res
}

func Part1(input []string) int {
	network := parseInput(input)
	start := "AAA"
	end := "ZZZ"

	element := start
	step := 0
	instructionsLenght := len(network.Instructions)
	for element != end {
		instruction := network.Instructions[step%instructionsLenght]
		if instruction == 'L' {
			element = network.Nodes[element][0]
		} else {
			element = network.Nodes[element][1]
		}
		step++
	}

	return step
} //zero divided by any number leaves a remainder of zero.
// 1%5 = 1 When a smaller number is divided by a larger one, the remainder is the smaller number itself

func Part2(input []string) int {
	network := parseInput(input)

	starts := []string{}
	for node := range network.Nodes {
		lastIndex := len(node) - 1
		if node[lastIndex] == 'A' {
			starts = append(starts, node)
		}
	}

	steps := make([]int, len(starts))
	instructionsLenght := len(network.Instructions)
	for i := 0; i < len(starts); i++ {
		element := starts[i]
		lastIndex := len(element) - 1
		for element[lastIndex] != 'Z' {
			instruction := network.Instructions[steps[i]%instructionsLenght]
			if instruction == 'L' {
				element = network.Nodes[element][0]
			} else {
				element = network.Nodes[element][1]
			}
			steps[i]++
		}
	}

	res := lcmSlice(steps)
	return res
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
	fileName := "../test_input"
	input := ReadFile(fileName)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input))
	fmt.Println(time.Since(start2))
}
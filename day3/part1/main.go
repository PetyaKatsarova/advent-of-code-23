
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../test_input")
	// file, err := os.Open("../puzzle_input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, val := range grid {
		fmt.Println(val)
	}

	fmt.Println("Sum of part numbers:", sumPartNumbers(grid))
}

func isSymbol(c byte) bool {
	return c != '.' && !('0' <= c && c <= '9')
}

func isAdjacentToSymbol(grid []string, x, y int) bool {
	gridX := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	gridY := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < 8; i++ {
		numX, numY := x+gridX[i], y+gridY[i]
		if numX >= 0 && numX < len(grid) && numY >= 0 && numY < len(grid[numX]) && isSymbol(grid[numX][numY]) {
			return true
		}
	}
	return false
}

func sumPartNumbers(grid []string) int {
	sum := 0
	for i, row := range grid {
		for j := 0; j < len(row); {
			c := row[j]
			if '0' <= c && c <= '9' {

				start := j
				// Find the end of the number
				for j < len(row) && '0' <= row[j] && row[j] <= '9' {
					j++
				}
				numStr := row[start:j]
				num, _ := strconv.Atoi(numStr)

				// Check if any digit of the number is adjacent to a symbol
				for k := start; k < j; k++ {
					if isAdjacentToSymbol(grid, i, k) {
						fmt.Println("Adding number:", num)
						sum += num
						break
					}
				}
			} else {
				j++
			}
		}
	}
	return sum
}

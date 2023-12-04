// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"regexp"
// 	"strconv"
// )

// func main() {
// 	fmt.Println("total is: ", getTotal())
// }

// /* case 1: symbol b4 and after
// 2. prev, next line same prev or next index with symbol

// */

// func getTotal() int {
// 	// file, err := os.Open("../puzzle_input")
// 	file, err := os.Open("../test_input")
// 	if err != nil {
// 		log.Fatalf("failed to open file: %s", err)
// 	}
// 	defer file.Close()

// 	total := 0
// 	// counter := 1 // for testing only
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		total += isAdjacentSymbol(line)
// 		// fmt.Printf("line %d: %s\n", counter, line)
// 		// counter++
// 	}
// 	return total
// }

// // checks if symbol b4 or after int and returns the sum
// func isAdjacentSymbol(line string) int {
// 	sum := 0
// 	// \d+ match 1 or more digits   [^\w\s.] matches any char exept word, white space or dot
// 	// This is where non-capturing groups come in, denoted by (?: ...). They group the included
// 	// pattern but do not capture the match for later use.
// 	genExp1 := regexp.MustCompile(`\d+[^\w\s.]`)
// 	genExp2 := regexp.MustCompile(`[^\w\s.]\d+`)
// 	matches := genExp1.FindAllString(line, -1)
// 	matches2 := genExp2.FindAllString(line, -1)
// 	for _, val := range matches {
// 			// Match sequences of digits possibly surrounded by non-word characters (excluding space and period)
// 		sum += extractInt(val, sum)

// 	}
// 	for _, m := range matches2 {
// 		sum += extractInt(m, sum)
// 		// re := regexp.MustCompile(`\d+`)
// 		// bla := re.FindAllString(m, -1)
// 		// for _, foo := range bla {
// 		// 	fmt.Println(foo)
// 		// }
// 	}
// 	return sum
// }

// func extractInt(val string, sum int) int {
// 	re := regexp.MustCompile(`\d+`)
// 	bla := re.FindAllString(val, -1)
// 	for _, val2 := range bla {
// 		num,_ := strconv.Atoi(val2)
// 		sum += num
// 		fmt.Printf("type of %d is %T\n", num, num)
// 	}
// 	return sum
// }

// // checks if symbol is on the same index or prev or next on prev or next row

// // make a slice of symbols

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("../test_input")
	file, err := os.Open("../puzzle_input")
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
				// Start of a number
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
						fmt.Println("Adding number:", num) // Print the number being added
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

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func isSymbol(c byte) bool {
    return c != '.' && !('0' <= c && c <= '9')
}

func getAdjacentNumbers(grid []string, x, y int) ([]int, error) {
    dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
    dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
    
    partNumbers := make(map[int]bool)
    
    for i := 0; i < 8; i++ {
        nx, ny := x+dx[i], y+dy[i]
        if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[nx]) && '0' <= grid[nx][ny] && grid[nx][ny] <= '9' {
            numStr := ""
            // Move left for multi-digit numbers
            for ny-1 >= 0 && '0' <= grid[nx][ny-1] && grid[nx][ny-1] <= '9' {
                ny--
            }
            // Capture the entire number
            for ny < len(grid[nx]) && '0' <= grid[nx][ny] && grid[nx][ny] <= '9' {
                numStr += string(grid[nx][ny])
                ny++
            }
            num, _ := strconv.Atoi(numStr)
            partNumbers[num] = true
        }
    }

    if len(partNumbers) != 2 {
        return nil, fmt.Errorf("not a valid gear at %d,%d", x, y)
    }

    var numbers []int
    for num := range partNumbers {
        numbers = append(numbers, num)
    }
    return numbers, nil
}

func sumGearRatios(grid []string) int {
    sum := 0
    for i, row := range grid {
        for j, c := range row {
            if c == '*' {
                nums, err := getAdjacentNumbers(grid, i, j)
                if err == nil {
                    sum += nums[0] * nums[1]
                }
            }
        }
    }
    return sum
}

func main() {
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

    fmt.Println("Sum of gear ratios:", sumGearRatios(grid))
}

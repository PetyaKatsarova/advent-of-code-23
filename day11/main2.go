// package main

// // learning from: https://github.com/OscarBrugne/AdventOfCode

// import (
// 	"fmt"
// 	"time"

// 	"main.go/utils"
// )

// // if free row: double it
// func doubleFreeRow(input []string) []string {
// 	if len(input) == 0 {
//         return input
//     }
// 	result := []string{}
// 	for _, val := range input { // loop through the lines of code
// 		isGalaxy := false
// 		// loop through the . and #
// 		for _, val2 := range val {
// 			if val2 == '#' { 
// 				isGalaxy = true
// 			}
// 		}
// 		result = append(result, val)
// 		if isGalaxy == false {
// 			result = append(result, val)
// 		}
// 	}
// 	return result
// }

// // doulbe free column
// func doubleFreeColumn(input []string) []string {
// 	if len(input) == 0 {
//         return input
//     }
// 	// find which cols need to be doubled
// 	doubleCol := make([]bool, len(input[0]))
// 	for i := range doubleCol {
// 		doubleCol[i] = true
// 		for _, row := range input {
// 			if row[i] == '#' {
// 				doubleCol[i] = false
// 				break
// 			}
// 		}
// 	}

// 	result := make([]string, len(input))
// 	for i, row := range input { // loop through rows
// 		var newRow string
// 		for j, ch := range row { // loop through cols
// 			newRow += string(ch)
// 			if doubleCol[j] {
// 				newRow += string(ch)
// 			}
// 		}
// 		result[i] = newRow
// 	}
// 	return result
// }

// type Point struct {
// 	x, y int
// }


// // sotres the coordinates of all the galaxies
// func findGalaxies(input []string) []Point {
// 	var galaxies []Point
// 	for y, row := range input {
// 		for x, col := range row {
// 			if col == '#' {
// 				galaxies = append(galaxies, Point{x,y})
// 			}
// 		}
// 	}
// 	return galaxies
// }



// // breath-first-search algorithm for traversing or searching ree or graph data struct
// func bfs(grid []string, start, end Point) int {
// 	queue := []Point{start}
// 	visited := make(map[Point]bool)
// 	visited[start] = true
// 	distance := 0
// 	dirs := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// 	for len(queue) > 0 {
// 		size := len(queue)
// 		for i := 0; i < size; i++ {
// 			cur := queue[0]
// 			queue = queue[1:]

// 			if cur == end {
// 				return distance
// 			}

// 			for _, dir := range dirs {
// 				next := Point{cur.x + dir.x, cur.y + dir.y}
// 				if next.x >= 0 && next.x < len(grid[0]) && next.y >= 0 && next.y < len(grid) && !visited[next] && grid[next.y][next.x] != '.' {
// 					visited[next] = true
// 					queue = append(queue, next)
// 				}	
// 			}
// 		}
// 		distance++
// 	}

// 	return -1
// }

// // func main() {
// // 	input := utils.ReadFile("test_input") // returns []string : per line
// // 	doubledRows := doubleFreeRow(input)
// // 	expendedUniverse	:= doubleFreeColumn(doubledRows)
// // 	for key, val := range expendedUniverse {
// // 		fmt.Println(key, ": ", val, "len: ", len(val))
// // 	}
// // 	galaxies := findGalaxies(expendedUniverse)
// // 	fmt.Println("num galaxies: ", len(galaxies))

// // 	startTime 	:= time.Now()
// // 	totalLength := 0
// // 	for i := 0; i < len(galaxies); i++ {
// // 		for j := i + 1; j < len(galaxies); j++ {
// // 			length := bfs(expendedUniverse, galaxies[i], galaxies[j])
// // 			totalLength += length
// // 		}
// // 	}
// // 	fmt.Println("result is: ", totalLength)
// // 	fmt.Println("processing time: ", time.Since(startTime))
// // }

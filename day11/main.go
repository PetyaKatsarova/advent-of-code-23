package main

import (
	"fmt"
	"time"

	"main.go/utils"
)

type Coord struct {
	X int
	Y int
}

type Grid struct {
	Width int
	Height int
	Data map[Coord]byte //grid stores only galaxies coord and the sing for it
}

var Empty byte = '.'

// grid keeps the coordinates of all the galaxies
func buildGrid(input []string, empty byte) Grid {
	grid := Grid {
		Width: len(input[0]),
		Height: len(input),
		Data: map[Coord]byte{},
	}

	for y, line := range input {
		for x, char := range line {
			if byte(char) != empty {
				grid.Data[Coord{x,y}] = byte(char) //  grid stores only galaxies, no empty
			}
		}
	}
	return grid
}

func (grid Grid) getEmptyCols() []int {
	emptyCols := []int{}
	for x:= 0; x < grid.Width; x++ { // nested slices length
		isEmpty := true
		y :=0
		for y < grid.Height { // first slice
			if _, ok := grid.Data[Coord{x,y}]; ok {
				isEmpty = false
			}
			y++
		}
		if isEmpty {
			emptyCols = append(emptyCols, x)
		}
	}
	return emptyCols
}

func (grid Grid) getEmptyRows() []int {
	emptyRows := []int{}

	for y :=0; y < grid.Height; y++ {
		isEmpty := true

		x := 0
		for x < grid.Width {
			if _, ok := grid.Data[Coord{x,y}]; ok {
				isEmpty = false
			}
			x++
		}
		if isEmpty {
			emptyRows = append(emptyRows, y)
		}
	}
	return emptyRows
}

func calculateOffsets(emptyIndexes []int, bound int) []int {
	offsets := make([]int, bound) // bound is len of the slice
	for _, idx := range emptyIndexes {
		for i := idx + 1; i < len(offsets); i++ {
			offsets[i]++
		}
	}
	return offsets
}

func expandGrid(grid Grid, expansionFactor int) Grid {
	emptyCols := grid.getEmptyCols()
	emptyRows := grid.getEmptyRows()
	numLinesToAdd := expansionFactor - 1

	newGrid := Grid {
		Width: grid.Width + len(emptyCols)*numLinesToAdd,
		Height: grid.Height + len(emptyRows)*numLinesToAdd,
		Data: make(map[Coord]byte, len(grid.Data)), //it is the same: coords of galaxies
	}

	dXs := calculateOffsets(emptyCols, grid.Width)
	dYs := calculateOffsets(emptyRows, grid.Height)

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			coord := Coord{X: x, Y: y}
			if _, ok := grid.Data[coord]; ok {
				newCoord := Coord{x + dXs[x]*numLinesToAdd, y + dYs[y]*numLinesToAdd}
				newGrid.Data[newCoord] = grid.Data[coord] // copies the byte sing for the new expanded coords
			}
		}
	}
	return newGrid
}



func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateLength(grid Grid, c1, c2 Coord) int {
	dX := abs(c2.X - c1.X)
	dY := abs(c2.Y - c1.Y)
	return dX + dY
}



func Part1(input []string) int {
	grid 			:= buildGrid(input, Empty)
	expandedGrid	:= expandGrid(grid, 2)
	res				:= 0
	visited			:= map[Coord]struct{}{}

	for coord1	:= range expandedGrid.Data { // loop through galaxies
		for corrd2 := range visited {
			length := calculateLength(expandedGrid, coord1, corrd2) // calc len of curr and visited
			res += length
		}
		visited[coord1] = struct{}{} // look at explanation bellow
	}

	return res
}

func Part2(input []string) int {
	grid 			:= buildGrid(input, Empty)
	expandedGrid	:= expandGrid(grid, 1000000)
	res				:= 0
	visited			:= map[Coord]struct{}{}

	for coord1	:= range expandedGrid.Data { // loop through galaxies
		for corrd2 := range visited {
			length := calculateLength(expandedGrid, coord1, corrd2) // calc len of curr and visited
			res += length
		}
		visited[coord1] = struct{}{} // look at explanation bellow
	}

	return res
}
/*
!! NB:
visited is a map where the keys (Coord values) represent the galaxies that have been visited.
 you don't need to store any associated values with these keys; you just need to check if a particular
  galaxy has been visited or not. So, instead of using map[Coord]bool, which would require storing a
   boolean value for each key, you can use map[Coord]struct{}.

struct{} is an empty struct type with no fields, and you create an instance of it using {}. 
It doesn't occupy any additional memory when used as the value type in a map. Therefore, you 
can use visited[coord1] = struct{}{} to add coord1 to the visited map as a key without any associated 
value. This is a common Go idiom for representing sets.
*/

func main() {
	fileName := "puzzle_input"
	input := utils.ReadFile(fileName)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input))
	fmt.Println(time.Since(start2))
}
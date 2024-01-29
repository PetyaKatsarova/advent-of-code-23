// learning from https://github.com/OscarBrugne/AdventOfCode

// read some theory at the bottom
package main

type Coord struct {
	X int
	Y int
}


func (c1 Coord) Add(c2 Coord) Coord {
	return Coord{c1.X + c2.X, c1.Y + c2.Y}
}

func (c1 Coord) Substract(c2 Coord) Coord {
	return Coord{c1.X - c2.X, c1.Y - c2.Y}
}

func (c Coord) Opposite() Coord {
	return Coord{-c.X, -c.Y}
}

type Grid struct {
	Width	int
	Height 	int
	Data	map[Coord]byte
}

const (
	Empty				 = '.' // byte for ASCII
	Start				 = 'S'
	Vertical        	 = '|'
	Horizontal			 = '-'
	BottomLeftCorner   	 = 'J' // the shape is: _| south-east: pointing to lowerleftcorner
	BottomRightCorner	 = 'L'
	TopLeftCorner 		 = '7'
	TopRightCorner 		 = 'F'
	Enclosed			 = 'X'
)	

var (
	Undefined	= Coord{0,0}
	Left		= Coord{0, -1}
    Bottom		= Coord{1, 0}
	Right 		= Coord{0, 1}
	Top			= Coord{-1, 0}
)

type Pipe map[Coord]struct{}

var CharToPipe = map[byte]Pipe {
	Vertical: 			Pipe{Top: struct{}{}, Bottom: struct{}{}}
	Horizontal:			Pipe{Left: struct{}{}, Right: struct{}{}},
	TopLeftCorner:		Pipe{Top: struct{}{}, Left: struct{}{}},
	TopRightCorner:     Pipe{Top: struct{}{}, Right: struct{}{}},
	BottomLeftCorner: 	Pipe{Bottom: struct{}{}, Left: struct{}{}},
	BottomRightCorner: 	Pipe{Bottom: struct{}{}, Right: struct{}{}},
}

func getPipeFromChar(char by) Pipe { // getByteFromMap[Coord]struct{}
	if pipe, ok := CharToPipe[Char]; ok { return pipe}
	return Pipe{}
}

func getCharFromPipe(pipe Pipe) Char {
	for char, associatedPipe := range CharToPipe {
		if pipe.isEqualPipe(associatedPipe) { return char }
	}
	return Empty
}

func (pipe1 Pipe) isEqualPipe(pipe2 Pipe) bool {
	if len(pipe1) != len(pipe2) { return false }
	for dir := range pipe1 {
		if _, ok := pipe2[dir]; !ok { return false }
	}
	return true
}

// type Grid struct {Width, height	int, data map[coord]Char}
func buildGrid(input []string) Grid {
	grid := Grid {
		Width: 	len(input[0]),
		Height: len(input),
		Data:   map[Coord]Char{},
	}

	for y, line := range input {
		for x, char := range line {
			if char != Empty {
				grid.Data[Coord{x, y}] = char
			}
		}
	}
	return grid
}

func (grid Grid) toString() string {
	pipesRepres := map[byte]string {
		Empty:             " ",
		Start:             "S",
		Vertical:          "║",
		Horizontal:        "═",
		TopLeftCorner:     "╝",
		TopRightCorner:    "╚",
		BottomLeftCorner:  "╗",
		BottomRightCorner: "╔",
		Enclosed:          "X",
	}

	var res string
	for y := 0; y < grid.Height; y++ {
		for x := 0; x <= grid.Width; x++ {
			coord := Coord{X: x, Y: y}
			if v, ok := grid.Data[coord]; ok {
				res += pipesRepres[v]
			} else {
				res += pipesRepres[Empty]
			}
		}
		res += "\n"
	}
	return res
}

func findStart(grid Grid) Coord {
	for coord, val := range grid.Data {
		if val == Start { return coord }
	}
	return Coord{}
}

// type Pipe map[Coord]struct {}
func (c Coord) getPipeFromNeighbours(grid Grid) Pipe {
	pipe := Pipe{}

	possibleNeighbours := map[Coord]Coord {
		Top:    c.Add(Top),
		Right:  c.Add(Right),
		Bottom: c.Add(Bottom),
		Left:   c.Add(Left),
	}

	for dir := range possibleNeighbours {
		if neighbourCoord, ok := possibleNeighbours[dir]
	}
	return pipe
}

/*
byte is used for representing 8-bit values, typically used for ASCII characters and binary data.
 rune is used for representing Unicode characters, which can be from different scripts and symbol sets.
  char is not a built-in type in Go, and characters are typically represented using byte or rune 
  depending on whether you're working with ASCII or Unicode characters.
*/
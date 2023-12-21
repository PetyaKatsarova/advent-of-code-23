package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"main.go/utils"
)

type Part struct {
	X int
	M int
	A int
	S int
}

func newPart(x, m, a, s int) Part {
	return Part{
		X: x,
		M: m,
		A: a,
		S: s,
	}
}

// 1. split on 2 arrs: returns workflow and parts as a slice of strings/lines
func getWorkflowAndParts(input []string) ([]string, []string) {
	ind := 0
	for i, val := range input {
		trimmedSpace := strings.TrimSpace(val)
		if trimmedSpace == "" {
			ind = i
		}
	}
	parts := input[ind+1:]
	cleanParts := []string{}
	for _, part := range parts {
		cleanParts = append(cleanParts, strings.Trim(part, "{}"))
	}
	return input[:ind], cleanParts
}

// returns a Part with all vals as ints
func getPartsVals(cleanParts string) Part {
	valuePairs := strings.Split(cleanParts, ",")
	// fmt.Println(len(valuePairs), ": ", valuePairs)
	if len(valuePairs) < 4 {
		panic("Not enough value pairs in the cleanParts")
	}

	x := convStrToInt(valuePairs[0])
	m := convStrToInt(valuePairs[1])
	a := convStrToInt(valuePairs[2])
	s := convStrToInt(valuePairs[3])

	part := newPart(x, m, a, s)
	return part
}

func convStrToInt(valuePair string) int {
	xStr := strings.Split(valuePair, "=")[1]
	xStr = strings.TrimSpace(xStr)
	// fmt.Println(len(xStr), ": ", xStr)
	x, err := strconv.Atoi(xStr)
	if err != nil {
		panic(err)
	}
	return x
}

// loop through all flow lines, start from in and follow the flow instructions with switch, checks for 1 part
// returns A,R for a part
func followTheFlow(lines []string, part Part) string {
	mapWorkFlow := map[string]string{} // {px: "a<2006:qkq,m>2090:A,rfg"} easy to work with

	for _, line := range lines {
		splittedLine	:= strings.Split(line, "{")
		key				:= splittedLine[0]
		val				:= splittedLine[1]
		val = strings.Trim(val, "}")
		mapWorkFlow[key] = val
	}

	nextLetters := processFlowLine("in", mapWorkFlow["in"], part)
	for nextLetters != "A" && nextLetters != "R" {
		nextLetters = processFlowLine(nextLetters, mapWorkFlow[nextLetters], part)
	}

	return nextLetters // should return only A || R
}

// returns next letters, R or A
func processFlowLine(nextLetters, line string, partStruct Part) string{
	// nextLetters = in, line = mapWorkFlow[nextLetters] = s<1351:px,qqz
	fmt.Println("line from processflowline", line)
	steps := strings.Split(line, ",")

	for _, val := range steps {
		if strings.Contains(val, ":") {
			splitVal := strings.Split(val, ":") // splitVal[0] = s<1351, splV[1] = px
			comparissonPair := splitVal[0]
			nextLetters		:= splitVal[1]
			partLetter		:= comparissonPair[:1] // only the first char s<1351
			sign			:= comparissonPair[1:2] // second char
		    numToCompare,err:= strconv.Atoi(string(comparissonPair[2:]))
			if err != nil { panic(err)}
			partValue		:= getLetterVal(partStruct, partLetter)
			if sign == "<" {
				if partValue < numToCompare {
					return nextLetters
				} else {
					continue
				}
			} else if sign == ">" {
				if partValue > numToCompare {
					return nextLetters
				} else {
					continue
				}
			}
		} else {
			return val
		}
	}
	return "WTF"
}

// return the int val of the corresponding Part field
func getLetterVal(part Part, partLetter string) int {
	partLetter = strings.ToUpper(partLetter)
	switch partLetter {
	case "X":
		return part.X //
	case "M":
		return part.M
	case "A":
		return part.A
	case "S":
		return part.S
	default:
		return -1
	}
}

func Part1(input []string) int {
	flow, parts := getWorkflowAndParts(input) // returns slices of strings
	var letter string
	sumResult := 0

	for _, part := range parts {
		structPart := getPartsVals(part) // returns Part
		letter = followTheFlow(flow, structPart) // 
		if letter == "A" {
			sumResult += structPart.X + structPart.M + structPart.A + structPart.S
		}
		fmt.Println("part1: ", letter)
	}
	// fmt.Println("flow: ",flow)
	return sumResult
}

func main() {
	input := utils.ReadFile("puzzle_input")

	start := time.Now()
	fmt.Println("result is: ", Part1(input))
	fmt.Println("time processing: ", time.Since(start))
}

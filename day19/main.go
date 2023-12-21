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
	if err != nil {	panic(err) }
	return x
}

// loop through flow lines, start from in and follow the flow instructions with switch
// returns A,R or? 
func followTheFlow(line []string, part Part) string {
	res := "L"
	for { // loop continuously till the result is A or R
		for _, val := range line {
			//need letters b4 {}, split on , follow instructions
			sth := strings.Split(val, "{")
			letters := sth[0]
			fmt.Println(letters)
		}
		if res == "A"
	}
	return res
}

// takes workflow []string from a line and returns 
// check beginning letters, follow inside the logic, return the result: R, A or?
func destructureFlowLine(flowLine []string) () {

}

func Part1(input []string) string {
	flow, parts := getWorkflowAndParts(input) // returns slices of strings 
	var letter string

	for _, part := range parts {
		structPart := getPartsVals(part)
		letter = followTheFlow(flow,structPart)
		fmt.Println("part1: ", structPart.A)
	}
	fmt.Println("flow: ",flow)
	return letter
}

func main() {
	input := utils.ReadFile("test_input")

	start := time.Now()
	fmt.Println("result is: ", Part1(input))
	fmt.Println("time processing: ", time.Since(start))
}

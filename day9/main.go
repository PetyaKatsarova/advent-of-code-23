package main

// learning from: https://github.com/OscarBrugne/AdventOfCode

import (
	"strconv"
	"strings"
	"time"
	"fmt"

	"main.go/utils"
)

func main() {
	fileName := "../puzzle_input"
	input := utils.ReadFile(fileName) // returns []str

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input))
	fmt.Println(time.Since(start2))
}

func Part1(input []string) int {
	histories	:= parseInput(input) // returns 2d slices of the original ints, instead of sl of strings
	res			:= 0

	for _, sl := range histories {
		extrapolationSeries := calculateExtrapolations(sl)

		futurePrediction := 0
		for i := len(extrapolationSeries)-1; i > -1; i-- {
			futurePrediction += extrapolationSeries[i][len(extrapolationSeries[i])-1]
		}
		res += futurePrediction
	}
	return res
}

// this was the toughest one for me
func calculateExtrapolations(history []int) [][]int {
	extrapolationsSeries := [][]int{}
	extrapolationsSeries = append(extrapolationsSeries, history)

	for i := 1; i < len(history); i++ {
		previousExptrapolation := extrapolationsSeries[i - 1]
		
		if isAllZero(previousExptrapolation) { return extrapolationsSeries }

		extrapolations := calculateExtrapolation(previousExptrapolation)
		extrapolationsSeries = append(extrapolationsSeries, extrapolations)
	}
	fmt.Println(extrapolationsSeries)
	return extrapolationsSeries
}

// extrapolation is: estimate data points outside of a set of observed data points.
func calculateExtrapolation(history []int) []int {
	resultsNums := []int{}

	for i := 0; i < len(history)-1; i++ {
		resultsNums = append(resultsNums, history[i+1] - history[i])
	}
	return resultsNums
}

func isAllZero(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			return false
		}
	}
	return true
}

func convertIntoInt(line string) []int {
	numsStr := strings.Fields(line)
	numsInt := []int{}

	for i := 0; i < len(numsStr); i++ {
		numInt, err := strconv.Atoi(numsStr[i])
		if err != nil {
			panic("procesLine func Atoi err: ")
		}
		numsInt = append(numsInt, numInt)
	}
	return numsInt
}

// returns 2d arr with all the int per line
func parseInput(input []string) (histories [][]int) {
	histories = [][]int{}
	for _, val := range input {
		nums := convertIntoInt(val)
		histories = append(histories, nums)
	}
	return histories
}

// part 2: 

func calculateLeftExtrapolation(history []int) []int {
	resultsNums := []int{}

	for i := 0; i < len(history)-1; i++ {
		resultsNums = append([]int{history[i+1] - history[i]}, resultsNums...) // prepend
	}
	return resultsNums
}

func calculateLeftExtrapolations(history []int) [][]int {
	extrapolationsSeries := [][]int{}
	extrapolationsSeries = append(extrapolationsSeries, history)

	for i := 1; i < len(history); i++ {
		previousExptrapolation := extrapolationsSeries[i - 1]
		
		if isAllZero(previousExptrapolation) { return extrapolationsSeries }

		extrapolations := calculateLeftExtrapolation(previousExptrapolation)
		extrapolationsSeries = append(extrapolationsSeries, extrapolations)
	}
	fmt.Println(extrapolationsSeries)
	return extrapolationsSeries
}

func Part2(input []string) int {
	histories	:= parseInput(input) // returns 2d slices of the original ints, instead of sl of strings
	res			:= 0

	for _, sl := range histories {
		extrapolationSeries := calculateExtrapolations(sl)

		pastPrediction := 0
		for i := len(extrapolationSeries)-1; i > -1; i-- {
			pastPrediction = extrapolationSeries[i][0] - pastPrediction
		}
		res += pastPrediction
	}
	return res
}
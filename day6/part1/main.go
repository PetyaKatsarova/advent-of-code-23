package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../puzzle_input")
	if err != nil {	log.Fatalf("failed to open file: %s", err) }
	defer file.Close()

	scanner	:= bufio.NewScanner(file)
	sth		:= make([]string, 0)// init empty slice
	for scanner.Scan() {
		line := scanner.Text()
		sth = append(sth, line)
	}
	fmt.Println("the result is: ", getResult(sth))
}

func getResult(sth []string) int {
	// WEIRD GO: IF THE SPLITTED STRING STARTS WITH THE SEPARATOR, THE FIRST EL OF THE SPLITTED
	// SLICE/ARR WILL BE ""
	time		:= strings.TrimSpace(strings.Split(sth[0], "Time: ")[1])
	distance	:=strings.TrimSpace(strings.Split(sth[1], "Distance: ")[1])
	timeArr		:= strings.Split(time, " ")
	distanceArr := strings.Split(distance, " ")
	trimmedTimeArr := trimSpaces(timeArr)
	trimmpedDistanceArr := trimSpaces(distanceArr)

	result		:= 1
	for i, val := range trimmedTimeArr {
		fmt.Println(val, trimmpedDistanceArr[i])
		bla,_ := strconv.Atoi(val)
		blabla,_ := strconv.Atoi(trimmpedDistanceArr[i])
		result *= calcRoundNum(bla, blabla )
	}
	return result
}

	// returns the number of possibilites to beat the record(distance)
func calcRoundNum(time int, distance int) int {
	count := 0
	for i := 0; i < time; i++ {
		hold := i
		// i milimeters per time - i
		traveledDistance := hold * ( time - i )
		if traveledDistance > distance {
			count++
		}
	}
	return count
}

func trimSpaces(arr []string) []string {
	result := make([]string, 0)
	for _, val := range arr {
		if val != "" {
			result = append(result, val)
		}
	}
	return result
}
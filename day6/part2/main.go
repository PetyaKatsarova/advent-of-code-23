package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
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
	fmt.Println("the result is: ", calcNum(sth))
	// bla, blu := getInts(sth)
	// fmt.Println("the result is: ", bla, blu)
}

func getInts(sth []string) (int, int) {
	timeStr := strings.TrimSpace(strings.Split(sth[0], "Time:")[1])
	distStr := strings.TrimSpace(strings.Split(sth[1], "Distance:")[1])
	fmt.Println(timeStr, distStr)
	timeStr	= strings.ReplaceAll(timeStr, " ", "")
	distStr	= strings.ReplaceAll(distStr, " ", "")
	time, err := strconv.Atoi(timeStr)
    if err != nil {
        log.Fatalf("failed to convert time to int: %s", err)
    }

    distance, err := strconv.Atoi(distStr)
    if err != nil {
        log.Fatalf("failed to convert distance to int: %s", err)
    }

    return time, distance
}

// 	// returns the number of possibilites to beat the record(distance)
func calcNum(line []string) int {
	time, distance := getInts(line)
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

// func trimSpaces(arr []string) []string {
// 	result := make([]string, 0)
// 	for _, val := range arr {
// 		if val != "" {
// 			result = append(result, val)
// 		}
// 	}
// 	return result
// }
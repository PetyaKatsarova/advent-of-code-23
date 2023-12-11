package main

// IT MELTS THE CPU, didnt optimize it: the result is: 9622622

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../puzzle_input")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	output := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}
	seeds, data := getSeeds(output)
	findTheLocation(createMap(data), seeds)
}

type Mapping struct {
	DestStart   int64
	SourceStart int64
	Length      int64
}

func createMap(data []string) map[string][]Mapping {
	result := map[string][]Mapping{}
	var currKey string

	for _, val := range data {
		if startWithAlpha(val) {
			currKey = val
			result[currKey] = []Mapping{}
		} else if currKey != "" {
			parts := strings.Fields(val) // fields splits on " "
			if len(parts) == 3 {
				destStart, _ := strconv.ParseInt(parts[0], 10, 64) //Atoi converts to int!
				sourceStart, _ := strconv.ParseInt(parts[1], 10, 64)
				length, _ := strconv.ParseInt(parts[2], 10, 64)
				result[currKey] = append(result[currKey], Mapping{DestStart: destStart, SourceStart: sourceStart, Length: length})
			}
		}
	}
	return result
}

// gets the seed and after manipulation returns the new number: it is per category
func applyMapping(mappings []Mapping, num int64) int64 {
	for _, mapping := range mappings {
		if num >= mapping.SourceStart && num < mapping.SourceStart+mapping.Length {
			offset := num - mapping.SourceStart
			return mapping.DestStart + offset
		}
	}
	return num
}

func findTheLocation(myMap map[string][]Mapping, seeds []int64) int64 {
	categories := []string{"seed-to-soil map:", "soil-to-fertilizer map:", "fertilizer-to-water map:", "water-to-light map:",
		"light-to-temperature map:", "temperature-to-humidity map:", "humidity-to-location map:"}
	minLocation := int64(math.MaxInt64)
	// transform the seeds:

	for _, seed := range seeds {
		currNum := seed
		for _, cat := range categories {
			mappings := myMap[cat] // gets the cats in the correct order !!
			currNum = applyMapping(mappings, currNum)
		}
		if currNum < minLocation {
			minLocation = currNum
		}
	}
	fmt.Println("the min location is: ", minLocation)
	return minLocation
}

func getSeeds(output []string) ([]int64, []string) {
	parts := strings.Split(output[0], "seeds:")
	if len(parts) < 2 {
		log.Fatalf("no seeds data found")
	}
	seedsStr := strings.Split(strings.TrimSpace(parts[1]), " ")
	seedsNmr := []int64{} // seeds in numbers
	data := output[1:]    // slice of all lines
	for i := 0; i < len(seedsStr); i += 2 {
		num, _ := strconv.ParseInt(seedsStr[i], 10, 64)
		count, _ := strconv.ParseInt(seedsStr[i+1], 10, 64)
		for j := num; j <= num+count; j++ {
			seedsNmr = append(seedsNmr, j)
		}

	}
	return seedsNmr, data
}

func startWithAlpha(s string) bool {
	if len(s) > 0 && s[0] >= 'a' && s[0] <= 'z' {
		return true
	}
	return false
}

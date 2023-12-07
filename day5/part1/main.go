package main

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
	if err != nil {	log.Fatalf("failed to open file: %s", err)}
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
	DestStart 	int
	SourceStart int
	Length		int
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
                destStart, _ := strconv.Atoi(parts[0])
                sourceStart, _ := strconv.Atoi(parts[1])
                length, _ := strconv.Atoi(parts[2])
                result[currKey] = append(result[currKey], Mapping{DestStart: destStart, SourceStart: sourceStart, Length: length})
            }
        }
    }
    return result
}

// gets the seed and after manipulation returns the new number: it is per category
func applyMapping(mappings []Mapping, num int) int {
	for _, mapping := range mappings {
		if num >= mapping.SourceStart && num < mapping.SourceStart + mapping.Length {
			offset := num - mapping.SourceStart
			return mapping.DestStart + offset
		}
	}
	return num
}

func findTheLocation(myMap map[string][]Mapping, seeds []int) int {
	categories := []string{"seed-to-soil map:", "soil-to-fertilizer map:", "fertilizer-to-water map:", "water-to-light map:",
	 "light-to-temperature map:", "temperature-to-humidity map:", "humidity-to-location map:"}
	 minLocation := math.MaxInt32

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

func getSeeds(output []string) ([]int, []string) {
	parts := strings.Split(output[0], "seeds:")
    if len(parts) < 2 {
        log.Fatalf("no seeds data found")
    }
    seedsStr := strings.Split(strings.TrimSpace(parts[1]), " ")
	seedsNmr := []int{} // seeds in numbers
	data	 := output[1:] // slice of all lines
	for _, val := range seedsStr {
		num,_ := strconv.Atoi(string(val))
		seedsNmr = append(seedsNmr, num)
	}
	return seedsNmr, data
}

func startWithAlpha(s string) bool {
	if len(s) > 0 && s[0] >= 'a' && s[0] <= 'z' {
		return true
	}
	return false
}


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
	fmt.Println("total is:", calcTotal())
}

func calcTotal() int {
	file, err := os.Open("../puzzle_input")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		games := strings.Split(line, ":")
		gameTitle := games[0]
		gameTries := games[1]
		isVald := true
		// fmt.Println(games[0]," ******** ", gameTries)
		tries := strings.Split(gameTries, ";")
		for _, val := range tries {
			opit := strings.Split(val, ", ")
			fmt.Println("-----------------------------------------------")
			for _, val := range opit {
				if checkColorNum(val) == false {
					isVald = false
				}			
			}
		}		
		if isVald == true {
			strNum := strings.Split(gameTitle, "Game ")
			bla := strings.TrimSpace(strNum[1])
			fmt.Println("-- game num --", bla)
			blabal, _ := strconv.Atoi(bla)
			fmt.Printf("type of strnum is: %T\n", blabal)
			total += blabal // calc the games[0].split at " " part [1]
		}
	}
	return total
 }

func checkColorNum(try string) bool {
	try = strings.TrimSpace(try)
	splittedTry := strings.Split(try, " ")
	color := splittedTry[1]
	num, err := strconv.Atoi(splittedTry[0])
	fmt.Println(num, "---", color)
    if err != nil {
        fmt.Println("Error converting string to integer:", err)
        return false
    }

	switch color {
	case "red": if num > 12 {
		return false
	}
	case "green": if num > 13 {
		return false
	}
	case "blue" : if num > 14 {
		return false
	}
	default: return true
	}
	return true
}

//part1> Aho Corasick ook

package main

import (
	"reflect"
	"testing"
)

var fileNameTest1 string = "../test_input"
var inputTest1 []string = ReadFile(fileNameTest1)

var fileNameTest2 string = "../test_input_part2"
var inputTest2 []string = ReadFile(fileNameTest2)

var fileNameDay string = "../puzzle_input"
var inputDay []string = ReadFile(fileNameDay)

func TestPart1(t *testing.T) {
	result := Part1(inputTest1)
	expected := 2
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest2)
	expected := 6
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestParseInput(t *testing.T) {
	input := []string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (CCC, AAA)",
	}
	result := parseInput(input)
	expected := Network{
		Instructions: "RL",
		Nodes: map[string][2]string{
			"AAA": {"BBB", "CCC"},
			"BBB": {"CCC", "AAA"},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}


/*
this benchmark test is used to measure the performance of the Part1 function. By repeatedly running
 Part1 for different values of b.N, the Go testing framework can determine how long on average it
  takes to execute Part1. This is useful for identifying performance bottlenecks, comparing different
   implementations for efficiency, or checking if new changes in the codebase have impacted performance.
*/
func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(inputDay)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(inputDay)
	}
}
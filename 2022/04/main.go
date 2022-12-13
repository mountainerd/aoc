package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func doesOneOfTheseCompletelyOverlapTheOther(elfPair []int) bool {
	switch {
	case elfPair[0] >= elfPair[2] && elfPair[1] <= elfPair[3]:
		fallthrough
	case elfPair[0] <= elfPair[2] && elfPair[1] >= elfPair[3]:
		return true
	}

	return false
}

func doTheseOverlapAtAll(elfPair []int) bool {
	switch {
	case elfPair[0] >= elfPair[2] && elfPair[0] <= elfPair[3]:
		fallthrough
	case elfPair[1] >= elfPair[2] && elfPair[1] <= elfPair[3]:
		fallthrough
	case elfPair[2] >= elfPair[0] && elfPair[2] <= elfPair[1]:
		fallthrough
	case elfPair[3] >= elfPair[0] && elfPair[3] <= elfPair[1]:
		return true
	}

	return false
}

func getAllMyNumbersFromTheString(elfPair string) []int {
	re := regexp.MustCompile(`\d+`)
	assignments := re.FindAllString(elfPair, -1)

	var numbers []int
	for _, assignment := range assignments {
		convertedString, _ := strconv.Atoi(assignment)
		numbers = append(numbers, convertedString)
	}

	return numbers
}

func PartOne(alternateInput ...string) int {
	var totalOverlappingPairs int

	inputFile := "input"

	if len(alternateInput) > 0 {
		inputFile = alternateInput[0]
	}

	input, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		elfPair := getAllMyNumbersFromTheString(scanner.Text())

		if doesOneOfTheseCompletelyOverlapTheOther(elfPair) {
			totalOverlappingPairs++
		}
	}

	return totalOverlappingPairs
}

func PartTwo(alternateInput ...string) int {
	var totalAnyOverlap int

	inputFile := "input"

	if len(alternateInput) > 0 {
		inputFile = alternateInput[0]
	}

	input, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		elfPair := getAllMyNumbersFromTheString(scanner.Text())

		if doTheseOverlapAtAll(elfPair) {
			totalAnyOverlap++
		}
	}

	return totalAnyOverlap
}

func main() {
	println("Part One:", PartOne())
	println("Part Two:", PartTwo())
}

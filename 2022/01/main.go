package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var ElvenCalories string

func calculateElfLoads() []int {
	var totalLoads []int

	elfLoads := convertLoadsToIntegers()

	for _, elfLoad := range elfLoads {
		var totalLoad int
		for _, load := range elfLoad {
			totalLoad += load
		}

		totalLoads = append(totalLoads, totalLoad)
	}

	return totalLoads
}

func convertLoadsToIntegers() [][]int {
	var convertedLoads [][]int

	elfLoads := separateList()

	for _, elf := range elfLoads {
		var convertedLoad []int
		caloriePacks := strings.Split(elf, "\n")
		for _, caloriePack := range caloriePacks {
			convertedPack, _ := strconv.Atoi(caloriePack)
			convertedLoad = append(convertedLoad, convertedPack)
		}

		convertedLoads = append(convertedLoads, convertedLoad)
	}

	return convertedLoads
}

func findLargestLoad(allLoads []int) int {
	var largestLoad int

	for _, load := range allLoads {
		if load > largestLoad {
			largestLoad = load
		}
	}

	return largestLoad
}

func findTopThreeLoadsTotal(allLoads []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(allLoads)))

	return allLoads[0] + allLoads[1] + allLoads[2]
}

func separateList() []string {
	return strings.Split(ElvenCalories, "\n\n")
}

func PartOne() int {
	calculatedLoads := calculateElfLoads()
	largestLoad := findLargestLoad(calculatedLoads)

	return largestLoad
}

func PartTwo() int {
	calculatedLoads := calculateElfLoads()
	return findTopThreeLoadsTotal(calculatedLoads)
}

func main() {
	fmt.Println("Part One:", PartOne())
	fmt.Println("Part Two:", PartTwo())
}

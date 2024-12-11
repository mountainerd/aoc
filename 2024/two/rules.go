package main

import (
	"github.com/mountainerd/aoc/2024/utils"
	"slices"
)

func ReportProcessor(reports [][]int) [][]int {
	// base case
	if len(reports) == 0 {
		return [][]int{}
	}

	// make us some channels to do work in parallel
	diffIsBad := make(chan bool, 1)
	dirIsBad := make(chan bool, 1)

	// go func diff
	go func() {
		diffIsBad <- differentialIsBad(reports[0])
	}()

	// go func dir
	go func() {
		dirIsBad <- directionIsBad(reports[0])
	}()

	diffBad, dirBad := <-diffIsBad, <-dirIsBad

	if diffBad || dirBad {
		if isDampened := problemDampener(reports[0], 0); !isDampened {
			return ReportProcessor(reports[1:])
		}
	}

	// good response
	return append([][]int{reports[0]}, ReportProcessor(reports[1:])...)
}

// checks the absolute differential between values to ensure they are within 1 and 3.
// returns true if bad
// returns false/default if there are no problems
func differentialIsBad(report []int) bool {
	// base case
	if len(report) == 1 {
		return false
	}

	// calculate difference
	diff := utils.AbsDiff(report[0], report[1])

	// test
	if diff < 1 || diff > 3 {
		// bad response
		return true
	}

	// good response
	return differentialIsBad(report[1:])
}

// checks the slice to ensure the values are all moving in the same direction
// returns true if bad
// returns false/default if there are no problems
func directionIsBad(report []int) bool {
	// base case
	if len(report) == 2 {
		return false
	}

	// calculate
	next := report[1] - report[0]
	later := report[2] - report[1]

	// tests
	if (next < 0 && later > 0) || (next > 0 && later < 0) {
		return true
	}

	// good response
	return directionIsBad(report[1:])
}

func problemDampener(report []int, index int) bool {
	updatedReport := make([]int, len(report))
	copy(updatedReport, report)

	// base case
	if index >= len(report) {
		return false
	}

	// calculate
	if index == 0 {
		//updatedReport = report[1:]
		updatedReport = updatedReport[1:]
	} else if index == len(report)-1 {
		//updatedReport = report[:index]
		updatedReport = updatedReport[:index]
	} else {
		updatedReport = slices.Delete(updatedReport, index, index+1)
	}

	diffIsBad := differentialIsBad(updatedReport)
	dirIsBad := directionIsBad(updatedReport)

	// test
	if !diffIsBad && !dirIsBad {
		return true
	}

	return problemDampener(report, index+1)
}

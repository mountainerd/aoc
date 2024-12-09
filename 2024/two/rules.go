package main

import "github.com/mountainerd/aoc/2024/utils"

func ReportProcessor(reports [][]int) [][]int {
	if len(reports) == 0 {
		return [][]int{}
	}

	// make us some channels to do work in parallel
	diffIsBad := make(chan bool, 1)
	directionIsBad := make(chan bool, 1)

	// go func diff
	go func() {
		diffIsBad <- differentialCheck(reports[0])
	}()

	// go func dir
	go func() {
		directionIsBad <- directionalCheck(reports[0])
	}()

	switch {
	case !<-diffIsBad && !<-directionIsBad: // so we're checking here that BOTH are not bad (aka good)
		return append([][]int{reports[0]}, ReportProcessor(reports[1:])...)
	default:
		return ReportProcessor(reports[1:])
	}
}

// checks the absolute differential between values to ensure they are within 1 and 3.
// returns true if bad
// returns false/default if there are no problems
func differentialCheck(report []int) bool {
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
	return differentialCheck(report[1:])
}

// checks the slice to ensure the values are all moving in the same direction
// returns true if bad
// returns false/default if there are no problems
func directionalCheck(report []int) bool {
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
	return directionalCheck(report[1:])
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mountainerd/aoc/2024/utils"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

func createReports(f *os.File) ([][]int, error) {
	var reports [][]int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		reportValuesString := strings.Fields(scanner.Text())

		var reportValuesInt []int
		for _, v := range reportValuesString {
			intValue, err := strconv.Atoi(v)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Error converting value to int: %s", err.Error()))
			}

			reportValuesInt = append(reportValuesInt, intValue)
		}

		reports = append(reports, reportValuesInt)
	}

	return reports, nil
}

func allIncreasingOrDecreasing(reports *[][]int) error {
	for i := 0; i < len(*reports); {
		report := (*reports)[i]

		// bail out early if we're sorted
		if slices.IsSorted(report) {
			i++
			continue
		}

		// since only russ cox and god know why go doesn't implement a way to easily check if a slice is descending.
		slices.Reverse(report)
		if !slices.IsSorted(report) {
			*reports = slices.Delete(*reports, i, i+1)
		} else {
			// put things back the way i found it because aoc has a nasty habit of making me pay for that if i don't.
			slices.Reverse(report)
			i++
		}
	}

	return nil
}

func betweenOneAndThree(reports *[][]int) int {
	for i := 0; i < len(*reports); {
		deleteOccurred := false
		report := (*reports)[i]

		for idx := 0; idx < len(report)-1; {
			diff := utils.AbsDiff(report[idx], report[idx+1])

			if diff < 1 || diff > 3 {
				*reports = slices.Delete(*reports, i, i+1)
				deleteOccurred = true
				break
			}

			idx++
		}

		if !deleteOccurred {
			i++
		}
	}

	return len(*reports)
}

func main() {
	file, err := os.Open(os.Getenv("AOC_FILE"))
	if err != nil {
		slog.Error("error opening file", "error", err.Error())
		os.Exit(123)
	}
	defer file.Close()

	reports, err := createReports(file)
	if err != nil {
		slog.Error("error creating reports", "error", err.Error())
		os.Exit(1)
	}

	_ = allIncreasingOrDecreasing(&reports)

	fmt.Println("allegedly safe reports:", betweenOneAndThree(&reports))
}

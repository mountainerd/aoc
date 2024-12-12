package main

import (
	"fmt"
	_ "github.com/mountainerd/aoc/2024/utils"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.ReadFile(os.Getenv("AOC_FILE"))
	if err != nil {
		slog.Error("AOC_FILE not set")
		os.Exit(1)
	}

	fmt.Println("total:", calculateTotal(findPairs(string(input)), 0, true))
}

func findPairs(input string) [][]string {
	pattern := `(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`
	re := regexp.MustCompile(pattern)

	return re.FindAllStringSubmatch(input, -1)
}

func calculateTotal(pairs [][]string, total int, process bool) int {
	// base case
	if len(pairs) == 0 {
		return total
	}

	// processing switch
	switch pairs[0][0] {
	case "don't()":
		return calculateTotal(pairs[1:], total, false)
	case "do()":
		return calculateTotal(pairs[1:], total, true)
	}

	// do we process?
	if !process {
		return calculateTotal(pairs[1:], total, process)
	}

	// setup
	x, _ := strconv.Atoi(pairs[0][2])
	y, _ := strconv.Atoi(pairs[0][3])

	// happy path
	return calculateTotal(pairs[1:], x*y+total, process)
}

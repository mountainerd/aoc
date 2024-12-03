package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

// inputPath is the location of the provided file with the lists as inputs.
var inputPath string

// typically i am anti-init function however it is aoc and yolo.
func init() {
	_, exists := os.LookupEnv("AOC_INPUT_PATH")
	if !exists {
		slog.Error("AOC_INPUT_PATH environment variable not set")
		os.Exit(1)
	}

	inputPath = os.Getenv("AOC_INPUT_PATH")
}

// Siglocs holds the two lists and is where we hang some methods to do the dirty work.
type Siglocs struct {
	left, right []int
	input       string
	similarity  map[int]int
}

// ReadInput takes the values from the provided file and reads them into the appropriate slice.
func (s *Siglocs) ReadInput() bool {
	file, err := os.Open(inputPath)
	if err != nil {
		slog.Error("Failed to open input file")
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		leftValue, err := strconv.Atoi(values[0])
		if err != nil {
			slog.Error("Failed to parse left value", "error", err.Error())
			return false
		}
		rightValue, err := strconv.Atoi(values[1])
		if err != nil {
			slog.Error("Failed to parse right value", "error", err.Error())
			return false
		}

		s.left = append(s.left, leftValue)
		s.right = append(s.right, rightValue)

		s.similarity[rightValue] += 1
	}

	return true
}

// SortLists takes the two lists and sorts them lowest to highest, returning true if both are successful.
func (s *Siglocs) SortLists() bool {
	slices.Sort(s.left)
	slices.Sort(s.right)

	return slices.IsSorted(s.left) && slices.IsSorted(s.right)
}

// ComputeDistance just gets the absolute value between the two integers and adds them to the calculation.
func (s *Siglocs) ComputeDistance() int {
	var totalDistance int

	for i := range s.left {
		totalDistance += absDiff(s.left[i], s.right[i])
	}

	return totalDistance
}

// CalculateSimilarity ranges through the left list to see how often it occurs on the right, multiplies, and adds that
// to the overall tabulation.
func (s *Siglocs) CalculateSimilarity() int {
	var score int

	for _, value := range s.left {
		score += value * s.similarity[value]
	}

	return score
}

// NewSiglocs simply returns an object with the inputPath set.
func NewSiglocs() Siglocs {
	s := Siglocs{}

	s.input = inputPath
	s.similarity = make(map[int]int)

	return s
}

// absDiff implements a rudimentary absolute value function but for integers, not floats.
func absDiff(x, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}

func main() {
	s := NewSiglocs()

	// we read the file and update the two integer slices.
	ok := s.ReadInput()
	if !ok {
		slog.Error("Failed to read input file")
		os.Exit(1)
	}

	// we sort the lists.
	ok = s.SortLists()
	if !ok {
		slog.Error("Sorting the lists was unsuccessful.")
		os.Exit(1)
	}

	// now we compute the distance, similarity, and print.
	fmt.Println("distance:", s.ComputeDistance())
	fmt.Println("similarity:", s.CalculateSimilarity())
}

package utils

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

// AbsDiff implements a rudimentary absolute value function but for integers, not floats.
func AbsDiff[N int | float64](x, y N) N {
	if x < y {
		return y - x
	}

	return x - y
}

func ReadInputIntSlices(inputPath string) ([][]int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		slog.Error("Failed to open input file")
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var intSlices [][]int
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		var valuesInt []int
		for _, v := range values {
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

			valuesInt = append(valuesInt, i)
		}

		intSlices = append(intSlices, valuesInt)
	}

	return intSlices, nil
}

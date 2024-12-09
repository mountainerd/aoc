package main

import (
	"fmt"
	"github.com/mountainerd/aoc/2024/utils"
	"log/slog"
	"os"
)

var inputPath string

// yolo
func init() {
	_, exists := os.LookupEnv("AOC_FILE")
	if !exists {
		slog.Error("AOC_FILE environment variable not set")
		os.Exit(1)
	}

	inputPath = os.Getenv("AOC_FILE")
}

func main() {
	input, err := utils.ReadInputIntSlices(inputPath)
	if err != nil {
		slog.Error("Error reading input", "error", err.Error())
		os.Exit(1)
	}

	filtered := ReportProcessor(input)

	fmt.Println("length:", len(filtered))
}

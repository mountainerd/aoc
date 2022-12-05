package main_test

import (
	"embed"
	"testing"

	d2 "github.com/mountainerd/aoc/2022/02"
)

// Opponents moves:
//  A --> Rock
//  B --> Paper
//  C --> Scissors

// Response moves:
//  X --> Rock
//  Y --> Paper
//  Z --> Scissors

// Result scoring:
//  X --> Loss
//  Y --> Draw
//  Z --> Win

// Scoring:
//  Win --> 6 points
//  Draw --> 3 points
//  Loss --> 0 points
//   +
//  Rock --> 1 point
//  Paper --> 2 points
//  Scissors --> 3 points

// idea := map[string][string]int{
// 	 "A": {"X": 4, "Y": 2, "Z": 9},
// 	 "B": {"X": 7, "Y": 5, "Z": 3},
// 	 "C": {"X": 1, "Y": 8, "Z": 6},
// }

//go:embed test_inputs/*
var testfs embed.FS

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		testFile string
		want     int
	}{
		{
			name:     "provided example",
			testFile: "aoc_example",
			want:     15,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, _ := testfs.ReadFile("test_inputs/" + test.testFile + ".txt")
			d2.RPSCribSheet = string(data)

			if got := d2.PartOne(); got != test.want {
				t.Errorf("PartOne() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name     string
		testFile string
		want     int
	}{
		{
			name:     "provided example",
			testFile: "aoc_example",
			want:     12,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, _ := testfs.ReadFile("test_inputs/" + test.testFile + ".txt")
			d2.RPSCribSheet = string(data)

			if got := d2.PartTwo(); got != test.want {
				t.Errorf("PartTwo() = %v, want %v", got, test.want)
			}
		})
	}
}

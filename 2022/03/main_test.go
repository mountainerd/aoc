package main_test

import (
	_ "embed"
	"testing"

	d3 "github.com/mountainerd/aoc/2022/03"
)

//go:embed test_input.txt
var testInput string

func TestPartOne(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "provided example",
			want: 157,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d3.Rucksacks = testInput

			if got := d3.PartOne(); got != test.want {
				t.Errorf("PartOne() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "provided example",
			want: 70,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d3.Rucksacks = testInput

			if got := d3.PartTwo(); got != test.want {
				t.Errorf("PartTwo() = %v, want %v", got, test.want)
			}
		})
	}
}

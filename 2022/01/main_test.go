package main_test

import (
	"embed"
	"testing"

	d1 "github.com/mountainerd/aoc/2022/01"
)

//go:embed test_inputs/*
var testfs embed.FS

// TestDayOne returns the elf carrying the most calories.
func TestDayOne(t *testing.T) {
	tests := []struct {
		name      string
		test_file string
		want      int
	}{
		{
			name:      "three",
			test_file: "multiple_three",
			want:      3000,
		},
		{
			name:      "six",
			test_file: "multiple_six",
			want:      6000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, _ := testfs.ReadFile("test_inputs/" + test.test_file)
			d1.ElvenCalories = string(data)
			if got := d1.PartOne(); got != test.want {
				t.Errorf("DayOne() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestDayTwo(t *testing.T) {
	tests := []struct {
		name      string
		test_file string
		want      int
	}{
		{
			name:      "three",
			test_file: "multiple_three",
			want:      6000,
		},
		{
			name:      "six",
			test_file: "multiple_six",
			want:      15000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, _ := testfs.ReadFile("test_inputs/" + test.test_file)
			d1.ElvenCalories = string(data)
			if got := d1.PartTwo(); got != test.want {
				t.Errorf("DayTwo() = %v, want %v", got, test.want)
			}
		})
	}
}

package main_test

import (
	"os"
	"testing"

	d5 "github.com/mountainerd/aoc/2022/05"
)

var input = [][]string{
	{},
	{"Z", "N"},
	{"M", "C", "D"},
	{"P"},
}

func TestPartOne(t *testing.T) {
	orders, _ := os.Open("test_commands")
	defer orders.Close()

	input := d5.CargoCrane{
		Stacks: input,
		Moves:  orders,
	}

	if got, want := d5.PartOne(input), "CMZ"; got != want {
		t.Errorf("PartOne() = %s, want %s", got, want)
	}
}

// func TestPartTwo(t *testing.T) {
// 	orders, _ := os.Open("test_commands")
// 	defer orders.Close()

// 	input := d5.CargoCrane{
// 		Stacks: input,
// 		Moves:  orders,
// 	}

// 	if got, want := d5.PartTwo(input), "MCD"; got != want {
// 		t.Errorf("PartTwo() = %s, want %s", got, want)
// 	}
// }

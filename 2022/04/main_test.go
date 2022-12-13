package main_test

import (
	"testing"

	d4 "github.com/mountainerd/aoc/2022/04"
)

func TestPartOne(t *testing.T) {
	if got, want := d4.PartOne("test_input"), 2; got != want {
		t.Errorf("PartOne() = %d, want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	if got, want := d4.PartTwo("test_input"), 4; got != want {
		t.Errorf("PartTwo() = %d, want %d", got, want)
	}
}

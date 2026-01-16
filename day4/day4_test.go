package main

import (
	"testing"
)

const input string = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func Test_part1(t *testing.T) {
	expected := 13
	actual := Part1(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_part2(t *testing.T) {
	expected := 43
	actual := Part2(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

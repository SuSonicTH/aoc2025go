package main

import (
	"testing"
)

const input string = `987654321111111
811111111111119
234234234234278
818181911112111
`

func Test_part1(t *testing.T) {
	expected := 357
	actual := Part1(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_part2(t *testing.T) {
	expected := 3121910778619
	actual := Part2(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

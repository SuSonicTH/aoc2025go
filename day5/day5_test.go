package main

import (
	"testing"
)

const input string = `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func Test_part1(t *testing.T) {
	expected := 3
	actual := Part1(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_part2(t *testing.T) {
	expected := 14
	actual := Part2(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

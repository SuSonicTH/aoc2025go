package main

import (
	"testing"
)

const input string = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func Test_part1(t *testing.T) {
	expected := 3
	actual := Part1(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_part2(t *testing.T) {
	expected := 6
	actual := Part2(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_part2_R1000(t *testing.T) {
	expected := 10
	actual := Part2("R1000")

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_part2_L51(t *testing.T) {
	expected := 1
	actual := Part2("L51")

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

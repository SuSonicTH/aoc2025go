package main

import (
	"testing"
)

const input string = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`

func Test_part1(t *testing.T) {
	expected := 4277556
	actual := Part1(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_part2(t *testing.T) {
	expected := 3263827
	actual := Part2(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

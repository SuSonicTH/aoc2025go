package main

import (
	"testing"
)

const input string = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124
`

func Test_part1(t *testing.T) {
	expected := 1227775554
	actual := Part1(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_part2(t *testing.T) {
	expected := 4174379265
	actual := Part2(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_isvalid2(t *testing.T) {
	var validTests = []struct {
		val   int
		valid bool
	}{
		{99, true},
		{111, true},
		{110, false},
		{565656, true},
		{5656565, false},
		{56565650, false},
	}

	for _, tt := range validTests {
		expected := tt.valid
		actual := isValid2(tt.val)

		if expected != actual {
			t.Errorf("expected %t, got %t for %d", expected, actual, tt.val)
		}
	}
}

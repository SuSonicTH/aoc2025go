package main

import (
	"aoc2025/util"
	"fmt"
	"strconv"
	"strings"
)

const day = 1

func main() {
	fmt.Printf("running day %d\n", day)
	input := util.ReadInput(day)
	util.TimedRun(1, Part1, input)
	util.TimedRun(2, Part2, input)
}

func parse(input string) (rotations []int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		direction := 1
		if line[0] == 'L' {
			direction = -1
		}

		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			panic("could not parse input '" + line + "'")
		}

		rotations = append(rotations, rotation*direction)
	}
	return
}

func Part1(input string) int {
	rotations := parse(input)
	var position = 50
	var zeroes = 0
	for _, rot := range rotations {
		position = util.Mod(position+rot, 100)
		if position == 0 {
			zeroes++
		}
	}
	return zeroes
}

func Part2(input string) int {
	rotations := parse(input)

	var position int = 50
	var zeroes int = 0

	for _, rot := range rotations {
		if rot < 0 {
			if position == 0 {
				zeroes--
			}
			position += rot
			zeroes -= util.FloorDiv(position, 100)
			if util.Mod(position, 100) == 0 {
				zeroes++
			}
		} else {
			position += rot
			zeroes += util.FloorDiv(position, 100)

		}
		position = util.Mod(position, 100)
	}
	return zeroes
}

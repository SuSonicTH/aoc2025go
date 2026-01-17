package main

import (
	"aoc2025/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(input string) int {
	return solve(input, 2)
}

func Part2(input string) int {
	return solve(input, 12)
}

func solve(input string, battieries int) int {
	joltage := 0
	for _, bank := range strings.Split(input, "\n") {
		if len(bank) > 0 {
			joltage += maxJoltage(bank, battieries)
		}
	}
	return joltage
}

func maxJoltage(bank string, batteries int) int {
	last := 0
	joltage := 0

	for battery := range batteries {
		candidate := bank[last : last+1]
		for pos := last + 1; pos <= len(bank)-batteries+battery; pos++ {
			current := bank[pos : pos+1]
			if current > candidate {
				candidate = current
				last = pos
			}
		}
		val, err := strconv.Atoi(candidate)
		if err != nil {
			panic(err)
		}

		joltage += int(math.Pow10(batteries-battery-1)) * val
		last++
	}

	return joltage
}

func main() {
	fmt.Println("running day 3")
	input := util.ReadInput(3)
	util.TimedRun(1, Part1, input)
	util.TimedRun(2, Part2, input)
}

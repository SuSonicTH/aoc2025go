package main

import (
	Grid "aoc2025/grid"
	"aoc2025/util"
	"fmt"
)

func Part1(input string) int {
	grid := Grid.ReadGrid(input)

	beams := make([]bool, grid.Width)
	start := grid.FindInRow(0, 'S')
	beams[start] = true

	split := 0
	for y := 1; y < grid.Height-1; y++ {
		for x := range grid.Width {
			if grid.Get(x, y) == '^' && beams[x] {
				split++
				beams[x-1] = true
				beams[x+1] = true
				beams[x] = false
			}
		}
	}
	return split
}

func Part2(input string) int {
	grid := Grid.ReadGrid(input)

	beams := make([]int, grid.Width)
	start := grid.FindInRow(0, 'S')
	beams[start] = 1

	for y := 1; y < grid.Height-1; y++ {
		for x := range grid.Width {
			if grid.Get(x, y) == '^' {
				beams[x-1] += beams[x]
				beams[x+1] += beams[x]
				beams[x] = 0
			}
		}
	}

	timelines := 0
	for _, beam := range beams {
		timelines += beam
	}
	return timelines
}

func main() {
	fmt.Println("running day 7")
	input := util.ReadInput(7)
	util.TimedRun(1, Part1, input)
	util.TimedRun(2, Part2, input)
}

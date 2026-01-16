package main

import (
	Grid "aoc2025/grid"
	"aoc2025/util"
	"fmt"
)

func main() {
	fmt.Println("running day 4")
	input := util.ReadInput(4)
	util.TimedRun(1, Part1, input)
	util.TimedRun(2, Part2, input)
}

func Part1(input string) int {
	grid := Grid.ReadGrid(input)
	var count int = 0
	for y := range grid.Height {
		for x := range grid.Width {
			if grid.Get(x, y) == '@' && grid.CountAdjecent(x, y, '@') < 4 {
				count++
			}
		}
	}
	return count
}

type Pos struct {
	x int
	y int
}

func Part2(input string) int {
	grid := Grid.ReadGrid(input)
	var count int = 0

	for {
		candidates := []Pos{}
		for y := range grid.Height {
			for x := range grid.Width {
				if grid.Get(x, y) == '@' && grid.CountAdjecent(x, y, '@') < 4 {
					candidates = append(candidates, Pos{x, y})
				}
			}
		}
		if len(candidates) == 0 {
			return count
		}
		count += len(candidates)
		for _, pos := range candidates {
			grid.Set(pos.x, pos.y, 'x')
		}
	}
}

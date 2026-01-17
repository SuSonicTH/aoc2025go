package grid

import (
	"fmt"
	"strings"
)

type Grid[T comparable] struct {
	Width  int
	Height int
	data   []T
}

func NewGrid[T comparable](widht int, height int) *Grid[T] {
	grid := Grid[T]{
		Width:  widht,
		Height: height,
		data:   make([]T, widht*height),
	}
	return &grid
}

func ReadGrid(input string) *Grid[rune] {
	widht, height := getInputDimentions(input)
	grid := NewGrid[rune](widht, height)
	for y, line := range strings.Split(input, "\n") {
		for x, val := range line {
			grid.Set(x, y, val)
		}
	}
	return grid
}

func getInputDimentions(input string) (width int, height int) {
	lines := strings.Split(input, "\n")
	width = len(lines[0])
	height = len(lines)
	if input[len(input)-1] == '\n' {
		height--
	}
	return
}

func (grid *Grid[T]) BoundsCheck(x, y int) {
	if x < 0 {
		panic(fmt.Sprintf("Out of Bounds: x of %d is less than 0", x))
	}
	if y < 0 {
		panic(fmt.Sprintf("Out of Bounds: y of %d is less than 0", y))
	}
	if x >= grid.Height {
		panic(fmt.Sprintf("Out of Bounds: x of %d is greter than maximum %d", x, grid.Width-1))
	}
	if y >= grid.Height {
		panic(fmt.Sprintf("Out of Bounds: y of %d is greter than maximum %d", y, grid.Height-1))
	}
}

func (grid *Grid[T]) InBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < grid.Width && y < grid.Height
}

func (grid *Grid[T]) Get(x, y int) T {
	grid.BoundsCheck(x, y)
	return grid.data[y*grid.Width+x]
}

func (grid *Grid[T]) Set(x, y int, value T) *Grid[T] {
	grid.BoundsCheck(x, y)
	grid.data[y*grid.Width+x] = value
	return grid
}

func (grid *Grid[T]) Reset(value T) *Grid[T] {
	for i := range grid.Width * grid.Height {
		grid.data[i] = value
	}
	return grid
}

func (grid *Grid[rune]) PrintRune() *Grid[rune] {
	for y := range grid.Height {
		for x := range grid.Height {
			fmt.Print(rune(grid.data[y*grid.Width+x]))
		}
		fmt.Println("")
	}
	return grid
}

func (grid *Grid[T]) Print() *Grid[T] {
	for y := range grid.Height {
		for x := range grid.Height {
			fmt.Print(grid.data[y*grid.Width+x])
		}
		fmt.Println("")
	}
	return grid
}

func (grid *Grid[T]) CountAdjecent(x, y int, value T) int {
	count := 0
	for _, yv := range []int{-1, 0, +1} {
		for _, xv := range []int{-1, 0, +1} {
			if xv == 0 && yv == 0 {
				continue
			}
			if grid.InBounds(x+xv, y+yv) && grid.Get(x+xv, y+yv) == value {
				count++
			}
		}
	}
	return count
}

func (grid *Grid[T]) FindInRow(y int, value T) int {
	for x := range grid.Width {
		if grid.Get(x, y) == value {
			return x
		}
	}
	return -1
}

func (grid *Grid[T]) FindInColumn(x int, value T) int {
	for y := range grid.Height {
		if grid.Get(x, y) == value {
			return y
		}
	}
	return -1
}

package util

import (
	"fmt"
	"math"
	"os"
	"time"
)

type partFunc func(string) int

func ReadInput(day int) string {
	return ReadFile(fmt.Sprintf("../input/input%d.txt", day))
}

func ReadFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TimedRun(part int, fun partFunc, input string) {
	start := time.Now()
	result := fun(input)
	runtime := time.Since(start)

	fmt.Printf("  Part%d: %d took %s\n", part, result, runtime)
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func FloorDiv(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}

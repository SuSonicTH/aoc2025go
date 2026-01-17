package main

import (
	"aoc2025/util"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	from int
	to   int
}

func parse(input string) (ranges []Range) {
	for _, current := range strings.Split(input[:len(input)-1], ",") {
		items := strings.Split(current, "-")
		from, err := strconv.Atoi(items[0])
		if err != nil {
			panic(fmt.Sprintf("could not parse '%s'", items[0]))
		}
		to, err := strconv.Atoi(items[1])
		if err != nil {
			panic(fmt.Sprintf("could not parse '%s'", items[1]))
		}
		ranges = append(ranges, Range{
			from: from,
			to:   to,
		})
	}
	return
}

type validFunc func(int) bool

func Part1(input string) int {
	return solve(input, isValid1)
}

func Part2(input string) int {
	return solve(input, isValid2)
}

func solve(input string, isValid validFunc) int {
	ranges := parse(input)
	var sum int = 0

	for _, r := range ranges {
		for i := r.from; i <= r.to; i++ {
			if isValid(i) {
				sum += i
			}
		}
	}
	return sum
}

func isValid1(i int) bool {
	val := strconv.Itoa(i)
	if util.Mod(len(val), 2) != 0 {
		return false
	}
	mid := len(val) / 2
	if val[:mid] == val[mid:] {
		return true
	}
	return false
}

func isValid2(i int) bool {
	val := strconv.Itoa(i)
	for length := 1; length < len(val); length++ {
		if len(val)%length == 0 {
			if isSubSeqValid(val, length) {
				return true
			}
		}
	}
	return false
}

func isSubSeqValid(str string, seqLen int) bool {
	subSeq := len(str) / seqLen
	for seq := 1; seq < subSeq; seq++ {
		first := (seq - 1) * seqLen
		second := seq * seqLen
		if str[first:first+seqLen] != str[second:second+seqLen] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("running day 2")
	input := util.ReadInput(2)
	util.TimedRun(1, Part1, input)
	util.TimedRun(2, Part2, input)
}

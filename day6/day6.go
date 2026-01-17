package main

import (
	"aoc2025/util"
	"fmt"
	"strconv"
	"strings"
)

func parse(input string) (numbers []string, ops string) {
	lines := strings.Split(input, "\n")

	numbers = lines[:len(lines)-2]
	ops = lines[len(lines)-2]
	return
}

func calculate(calc *int, op byte, number int) {
	switch op {
	case '+':
		*calc += number
	case '*':
		*calc *= number
	default:
		panic(fmt.Sprintf("Unknown operator '%c' expecting + or *", op))
	}
}

func Part1(input string) int {
	numbers, ops := parse(input)

	result := 0
	next := 0
	for pos := 0; pos < len(numbers[0]); {
		calc := 0
		for n, current := range numbers {
			start := pos
			for current[start] == ' ' {
				start++
			}
			end := start + 1
			for end < len(current) && current[end] != ' ' {
				end++
			}
			if end > next {
				next = end
			}

			number, err := strconv.Atoi(current[start:end])
			if err != nil {
				panic(fmt.Sprintf("Could not parse '%s' as integer", current[start:end]))
			}
			if n == 0 {
				calc = number
			} else {
				calculate(&calc, ops[pos], number)
			}
		}
		result += calc
		pos = next + 1
	}
	return result
}

func Part2(input string) int {
	numbers, ops := parse(input)

	result := 0
	for pos := len(ops) - 1; pos > 0; {
		o := pos
		for ops[o] == ' ' {
			o--
		}

		fistNumber := true
		calc := 0
		for pos >= o {
			number := parseNumber(numbers, pos)
			if fistNumber {
				calc = number
				fistNumber = false
			} else {
				calculate(&calc, ops[o], number)
			}
			if pos == 0 {
				return result + calc
			}
			pos -= 1
		}
		result += calc
		pos -= 1
	}

	return result
}

func parseNumber(numbers []string, pos int) (number int) {
	for _, num := range numbers {
		if num[pos] >= '0' && num[pos] <= '9' {
			number *= 10
			number += int(num[pos] - '0')
		}
	}
	return
}

func main() {
	fmt.Println("running day 6")
	input := util.ReadInput(6)
	util.TimedRun(1, Part1, input)
	util.TimedRun(2, Part2, input)
}

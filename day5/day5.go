package main

import (
	"aoc2025/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	from    int
	to      int
	deleted bool
}

type Data struct {
	fresh    []Range
	products []int
}

func (data *Data) isFresh(id int) bool {
	for _, r := range data.fresh {
		if id >= r.from && id <= r.to {
			return true
		}
	}
	return false
}

func (data *Data) sortFresh() {
	sort.Slice(data.fresh, func(i, j int) bool {
		return data.fresh[i].from < data.fresh[j].from
	})
}

func parse(input string) *Data {
	data := Data{
		fresh:    make([]Range, 0),
		products: make([]int, 0),
	}

	for _, line := range strings.Split(input, "\n") {
		dash := strings.IndexRune(line, '-')
		if dash > 0 {
			from, err := strconv.Atoi(line[:dash])
			if err != nil {
				panic(fmt.Sprintf("Could not parse '%s' as integer", line[:dash]))
			}

			to, err := strconv.Atoi(line[dash+1:])
			if err != nil {
				panic(fmt.Sprintf("Could not parse '%s' as integer", line[:dash]))
			}
			data.fresh = append(data.fresh, Range{from: from, to: to})
		} else if len(line) > 0 {
			product, err := strconv.Atoi(line)
			if err != nil {
				panic(fmt.Sprintf("Could not parse '%s' as integer", line))
			}
			data.products = append(data.products, product)
		}
	}
	return &data
}

func Part1(input string) int {
	data := parse(input)
	count := 0
	for _, id := range data.products {
		if data.isFresh(id) {
			count++
		}
	}
	return count
}

func Part2(input string) int {
	data := parse(input)
	data.sortFresh()

	for i := range len(data.fresh) - 1 {
		current := data.fresh[i]
		n := i + 1
		next := data.fresh[n]
		if current.to >= next.from {
			data.fresh[n].from = min(current.from, next.from)
			data.fresh[n].to = max(current.to, next.to)
			data.fresh[i].deleted = true
		}
	}

	count := 0
	for _, r := range data.fresh {
		if !r.deleted {
			count += r.to - r.from + 1
		}
	}
	return count
}

func main() {
	fmt.Println("running day 5")
	input := util.ReadInput(5)
	util.TimedRun(1, Part1, input)
	util.TimedRun(2, Part2, input)
}

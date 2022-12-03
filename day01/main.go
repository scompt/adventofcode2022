package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input2
var input2 string

func partOne(input string) int {
	parsed := parseInput(input)
	max := findMax(parsed)
	return max
}

func findMax(parsed [][]int) int {
	sums := make([]int, len(parsed))
	for i, perElf := range parsed {
		// fmt.Printf("%d %v\n", i, perElf)
		sums[i] = sumIntSlice(perElf)
	}
	return maxIntSlice(sums)
}

func sumIntSlice(ints []int) (m int) {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func maxIntSlice(v []int) (m int) {
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
		}
	}
	return
}

func parseInput(input string) [][]int {
	elves := make([][]int, 1)
	elves[0] = make([]int, 0)

	currentElf := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			currentElf += 1
			elves = append(elves, make([]int, 0))
			continue
		}

		line, err := strconv.Atoi(strings.TrimRight(line, "\n"))
		if err != nil {
			panic(err)
		}

		elves[currentElf] = append(elves[currentElf], line)
		// fmt.Println(line)
	}

	if len(elves[currentElf]) == 0 {
		elves = elves[:len(elves)-1]
	}

	return elves
}

func main() {
	output := partOne(input2)
	fmt.Println(output)
}

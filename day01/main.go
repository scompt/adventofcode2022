package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"scompt.com/adventofcode/twentytwentytwo/util"
)

//go:embed input2
var input2 string

func partTwo(input string) int {
	parsed := parseInput(input)
	maxes := findMax(parsed, 3)
	sum := util.SumIntSlice(maxes)
	return sum
}

func partOne(input string) int {
	parsed := parseInput(input)
	max := findMax(parsed, 1)
	return max[0]
}

func findMax(parsed [][]int, topN int) []int {
	maxes := make([]int, topN)
	sums := make([]int, len(parsed))
	for i, perElf := range parsed {
		sums[i] = util.SumIntSlice(perElf)
	}

	for i := 0; i < topN; i++ {
		// fmt.Printf("%d %v %v\n", i, maxes, sums)
		max, maxI := maxIntSlice(sums)
		maxes[i] = max
		sums[maxI] = 0
	}

	// return maxIntSlice(sums)
	return maxes
}

func maxIntSlice(v []int) (m int, maxI int) {
	if len(v) > 0 {
		m = v[0]
		maxI = 0
	}
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
			maxI = i
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

	output2 := partTwo(input2)
	fmt.Println(output2)
}

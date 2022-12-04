package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input2
var input2 string

func partOne(input string) (fullyContained int) {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")

	for _, line := range lines {
		left, right, _ := strings.Cut(line, ",")
		leftFromStr, leftToStr, _ := strings.Cut(left, "-")
		rightFromStr, rightToStr, _ := strings.Cut(right, "-")

		// fmt.Printf("%s %s %s %s\n", leftFromStr, leftToStr, rightFromStr, rightToStr)

		leftFrom, _ := strconv.Atoi(leftFromStr)
		leftTo, _ := strconv.Atoi(leftToStr)
		rightFrom, _ := strconv.Atoi(rightFromStr)
		rightTo, _ := strconv.Atoi(rightToStr)

		// fmt.Printf("%d %d %d %d\n", leftFrom, leftTo, rightFrom, rightTo)

		if (leftFrom >= rightFrom && leftTo <= rightTo) ||
			(rightFrom >= leftFrom && rightTo <= leftTo) {
			fullyContained += 1
		}
	}

	return
}

func main() {
	score1 := partOne(input2)
	fmt.Printf("%d\n", score1)
}

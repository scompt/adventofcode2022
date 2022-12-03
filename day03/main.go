package main

import (
	_ "embed"
	"fmt"
	"strings"

	"scompt.com/adventofcode/twentytwentytwo/util"
)

//go:embed input2
var input2 string

type rucksack struct {
	first  string
	second string
	size   int
}

func (s rucksack) findDupe() (r rune) {
	seen := map[rune]bool{}

	for _, v := range s.first {
		seen[v] = true
	}

	for _, r = range s.second {
		_, found := seen[r]
		if found {
			return
		}
	}
	panic("Not found")
}

func scoreChar(input rune) int {
	if strings.ToUpper(string(input)) == string(input) {
		return int(input) - 38
	} else {
		return int(input) - 96
	}
}

func parseInput(input string) []rucksack {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	rucksacks := make([]rucksack, len(lines))

	for i, line := range lines {
		rucksackSize := int(len(line) / 2)
		first := line[:rucksackSize]
		second := line[rucksackSize:]
		rucksacks[i] = rucksack{first, second, rucksackSize}
	}
	return rucksacks
}

func findAllDupes(rucksacks []rucksack) []rune {
	dupes := make([]rune, len(rucksacks))

	for i, rucksack := range rucksacks {
		dupes[i] = rucksack.findDupe()
	}

	return dupes
}

func scoreDupes(runes []rune) []int {
	scores := make([]int, len(runes))

	for i, dupe := range runes {
		scores[i] = scoreChar(dupe)
	}

	return scores
}

func partOne(input string) int {
	rucksacks := parseInput(input)
	// fmt.Printf("%v\n", rucksacks)
	dupes := findAllDupes(rucksacks)
	// fmt.Printf("%v\n", dupes)
	scores := scoreDupes(dupes)
	sum := util.SumIntSlice(scores)
	return sum
}

func main() {
	score1 := partOne(input2)
	fmt.Printf("%d\n", score1)
}

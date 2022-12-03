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
	all    string
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
		rucksacks[i] = rucksack{first, second, line}
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

func findBadge(group []string) rune {
	seen := map[rune]int{}
	for _, blah := range group {
		for _, v := range blah {
			value := seen[v]
			value += 1
			seen[v] = value
		}
	}
	// fmt.Printf("seen %v\n", seen)

	for r := range seen {
		if seen[r] == 3 {
			return r
		}
	}
	panic("whatwhat")
}

func findBadges(allStrings []string) []rune {
	var badges []rune

	for i := 0; i < len(allStrings); i += 3 {
		group := allStrings[i : i+3]
		// fmt.Printf("%v\n", group)
		badge := findBadge(group)
		// fmt.Printf("%v\n", string(badge))
		badges = append(badges, badge)
	}

	return badges
}

func removeDupes(rucksacks []rucksack) []string {
	unduped := make([]string, len(rucksacks))

	for i, r := range rucksacks {
		seen := map[rune]bool{}
		for _, aRune := range r.all {
			seen[aRune] = true
		}

		keys := make([]rune, 0, len(seen))
		for k := range seen {
			keys = append(keys, k)
		}
		unduped[i] = string(keys)
	}
	return unduped
}

func partTwo(input string) int {
	rucksacks := parseInput(input)
	// fmt.Printf("%v\n", rucksacks)
	unduped := removeDupes(rucksacks)
	// fmt.Printf("%v\n", unduped)
	badges := findBadges(unduped)
	// fmt.Printf("%v\n", badges)
	scores := scoreDupes(badges)
	sum := util.SumIntSlice(scores)
	return sum
}

func main() {
	score1 := partOne(input2)
	fmt.Printf("%d\n", score1)

	score2 := partTwo(input2)
	fmt.Printf("%d\n", score2)
}

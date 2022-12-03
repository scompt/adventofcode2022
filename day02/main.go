package main

import (
	_ "embed"
	"fmt"
	"strings"

	"scompt.com/adventofcode/twentytwentytwo/util"
)

//go:embed input2
var input2 string

type move struct {
	oppo string
	my   string
}

func (m move) score() int {
	return m.scoreMove() + m.scoreMyShape()
}
func (m move) scoreMyShape() int {
	if m.my == "X" {
		return 1
	} else if m.my == "Y" {
		return 2
	} else if m.my == "Z" {
		return 3
	}
	panic(m.my)
}
func (m move) scoreMove() int {
	if (m.oppo == "A" && m.my == "X") ||
		(m.oppo == "B" && m.my == "Y") ||
		(m.oppo == "C" && m.my == "Z") {
		return 3
	} else if (m.oppo == "A" && m.my == "Z") ||
		(m.oppo == "B" && m.my == "X") ||
		(m.oppo == "C" && m.my == "Y") {
		return 0
	} else {
		return 6
	}
}

func scoreMoves(moves []move) []int {
	scores := make([]int, len(moves))

	for i, m := range moves {
		scores[i] = m.score()
	}

	return scores
}

func parseInput(input string) []move {
	lines := strings.Split(input, "\n")
	var moves []move

	for _, line := range lines {
		if len(line) < 3 {
			break
		}

		oppoMove := string(line[0])
		myMove := string(line[2])
		moves = append(moves, move{oppoMove, myMove})
	}

	return moves
}

func partOne(input string) int {
	moves := parseInput(input)
	scores := scoreMoves(moves)
	// fmt.Printf("%v\n", scores)
	sum := util.SumIntSlice(scores)
	return sum
}
func main() {
	score := partOne(input2)
	fmt.Printf("%d\n", score)
}

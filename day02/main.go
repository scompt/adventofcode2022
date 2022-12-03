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
type move2 struct {
	oppo    string
	outcome string
}

func (m move2) score() int {
	return m.scoreOutcome() + m.scoreMyMove()
}
func (m move2) scoreOutcome() int {
	if m.outcome == "X" {
		return 0
	} else if m.outcome == "Y" {
		return 3
	} else if m.outcome == "Z" {
		return 6
	}
	panic(m.outcome)
}
func (m move2) scoreMyMove() int {
	var myShape string
	if m.outcome == "Y" {
		myShape = m.oppo
	} else if m.outcome == "X" {
		if m.oppo == "A" {
			myShape = "C"
		} else if m.oppo == "B" {
			myShape = "A"
		} else {
			myShape = "B"
		}
	} else {
		if m.oppo == "A" {
			myShape = "B"
		} else if m.oppo == "B" {
			myShape = "C"
		} else {
			myShape = "A"
		}
	}

	return scoreMyShape(myShape)
}

func (m move) score() int {
	return m.scoreMove() + scoreMyShape(m.my)
}
func scoreMyShape(shape string) int {
	if shape == "X" || shape == "A" {
		return 1
	} else if shape == "Y" || shape == "B" {
		return 2
	} else if shape == "Z" || shape == "C" {
		return 3
	}
	panic(shape)
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

func scoreMoves1(moves []move) []int {
	scores := make([]int, len(moves))

	for i, m := range moves {
		scores[i] = m.score()
	}

	return scores
}
func scoreMoves2(moves []move2) []int {
	scores := make([]int, len(moves))

	for i, m := range moves {
		scores[i] = m.score()
	}

	return scores
}

func parseInput1(input string) []move {
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
func parseInput2(input string) []move2 {
	lines := strings.Split(input, "\n")
	var moves []move2

	for _, line := range lines {
		if len(line) < 3 {
			break
		}

		oppoMove := string(line[0])
		outcome := string(line[2])
		moves = append(moves, move2{oppoMove, outcome})
	}

	return moves
}

func partOne(input string) int {
	moves := parseInput1(input)
	scores := scoreMoves1(moves)
	// fmt.Printf("%v\n", scores)
	sum := util.SumIntSlice(scores)
	return sum
}
func partTwo(input string) int {
	moves := parseInput2(input)
	scores := scoreMoves2(moves)
	// fmt.Printf("%v\n", scores)
	sum := util.SumIntSlice(scores)
	return sum
}
func main() {
	score1 := partOne(input2)
	fmt.Printf("%d\n", score1)

	score2 := partTwo(input2)
	fmt.Printf("%d\n", score2)
}

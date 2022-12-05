package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var RPSCribSheet string

var pointsReactionBased = map[string]map[string]int{
	"A": {"X": 4, "Y": 8, "Z": 3},
	"B": {"X": 1, "Y": 5, "Z": 9},
	"C": {"X": 7, "Y": 2, "Z": 6},
}

var pointsResultBased = map[string]map[string]int{
	"A": {"X": 3, "Y": 4, "Z": 8},
	"B": {"X": 1, "Y": 5, "Z": 9},
	"C": {"X": 2, "Y": 6, "Z": 7},
}

func tabulateScore(matchList []string, scoresheet map[string]map[string]int) int {
	var totalPoints int

	for _, match := range matchList {
		roundMoves := strings.Split(match, " ")
		opponentMove := roundMoves[0]
		responseMove := roundMoves[1]

		//totalPoints += pointsReactionBased[opponentMove][responseMove]
		totalPoints += scoresheet[opponentMove][responseMove]
	}

	return totalPoints
}

func splitIntoMatches() []string {
	return strings.Split(RPSCribSheet, "\n")
}

func PartOne() int {
	return tabulateScore(splitIntoMatches(), pointsReactionBased)
}

func PartTwo() int {
	return tabulateScore(splitIntoMatches(), pointsResultBased)
}

func main() {
	fmt.Println("Part One:", PartOne())
	fmt.Println("Part Two:", PartTwo())
}

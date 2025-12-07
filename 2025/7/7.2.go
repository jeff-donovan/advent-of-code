package main

import "strings"

type Coord struct {
	rowIndex int
	colIndex int
}

func travelToNextRow(lines []string, currentPosition Coord) []Coord {
	nextRow := lines[currentPosition.rowIndex+1]
	nextRowChar := string(nextRow[currentPosition.colIndex])
	// case 1 - next row is a dot - can only move down
	if nextRowChar == "." {
		return []Coord{{rowIndex: currentPosition.rowIndex + 1, colIndex: currentPosition.colIndex}}
	}

	// case 2 - next row is a splitter - can move left or right
	var legalMoves []Coord
	leftCoord := Coord{currentPosition.rowIndex + 1, currentPosition.colIndex - 1}
	rightCoord := Coord{currentPosition.rowIndex + 1, currentPosition.colIndex + 1}
	if isLegalCoord(lines, leftCoord) {
		legalMoves = append(legalMoves, leftCoord)
	}
	if isLegalCoord(lines, rightCoord) {
		legalMoves = append(legalMoves, rightCoord)
	}
	return legalMoves
}

func isLegalCoord(lines []string, c Coord) bool {
	return 0 <= c.rowIndex && c.rowIndex < len(lines) && 0 <= c.colIndex && c.colIndex < len(lines[0])
}

func algorithm7_2(lines []string) int {
	// plan
	//  - make a map of all coords to scores
	//  - i think give every dot in the last row a score of 1
	//  - loop backwards from bottom row to top
	//    - if it's a ^ then its score is 0
	//    - if it's a dot, use travelToNextRow() to figure out which spots it could travel to in the next row
	//    - score for the coord is the sum of the scores for the legal moves it could travel to
	//  - work all the way back up to the first row
	//  - return the score for the "S"

	scores := make(map[Coord]int)

	// pre-populate with last row
	for i := 0; i < len(lines[0]); i++ {
		scores[Coord{len(lines) - 1, i}] = 1
	}

	// loop backwards from bottom to top, skip bottom row because it was already pre-populated
	for rowIndex := len(lines) - 2; rowIndex >= 0; rowIndex-- {
		for i, char := range lines[rowIndex] {
			coord := Coord{rowIndex, i}
			if string(char) == "^" {
				scores[coord] = 0
				continue
			}

			score := 0
			for _, c := range travelToNextRow(lines, coord) {
				score += scores[c]
			}
			scores[coord] = score
		}
	}

	startIndex := strings.Index(lines[0], "S")
	return scores[Coord{0, startIndex}]
}

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
	startIndex := strings.Index(lines[0], "S")
	timelines := []Coord{{0, startIndex}}

	for rowIndex := 0; rowIndex < len(lines)-1; rowIndex++ {
		var newTimelines []Coord
		for _, c := range timelines {
			newTimelines = append(newTimelines, travelToNextRow(lines, c)...)
		}
		timelines = newTimelines
	}

	return len(timelines)
}

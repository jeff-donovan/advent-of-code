package main

import (
	"slices"
	"strings"
)

func updateNextGridRow(grid []string, currentRowIndex int) []string {
	var newGrid []string

	newGrid = append(newGrid, grid...)

	if currentRowIndex == 0 {
		startIndex := strings.Index(grid[0], "S")
		newGrid[1] = newGrid[1][0:startIndex] + "|" + newGrid[1][startIndex+1:]
		return newGrid
	}

	nextRowSplitters := findIndexesOfSplittersForRow(grid, currentRowIndex+1)
	pipeIndexes := findIndexesOfPipesForRow(grid, currentRowIndex)

	for _, pipeIndex := range pipeIndexes {
		nextRow := strings.Clone(newGrid[currentRowIndex+1])
		// first handle a continued pipe
		if !slices.Contains(nextRowSplitters, pipeIndex) {
			newGrid[currentRowIndex+1] = nextRow[0:pipeIndex] + "|" + nextRow[pipeIndex+1:]
		} else {
			newGrid[currentRowIndex+1] = nextRow[0:pipeIndex-1] + "|^|" + nextRow[pipeIndex+2:]
		}
	}

	return newGrid
}

func findIndexesOfPipesForRow(grid []string, rowIndex int) []int {
	var pipeIndexes []int
	for i, char := range grid[rowIndex] {
		if string(char) == "|" {
			pipeIndexes = append(pipeIndexes, i)
		}
	}
	return pipeIndexes
}

func findIndexesOfSplittersForRow(grid []string, rowIndex int) []int {
	if rowIndex >= len(grid) {
		return nil
	}

	var splitterIndexes []int
	for i, char := range grid[rowIndex] {
		if string(char) == "^" {
			splitterIndexes = append(splitterIndexes, i)
		}
	}
	return splitterIndexes
}

func algorithm7_1(lines []string) int {
	total := 0
	for rowIndex := 0; rowIndex < len(lines)-1; rowIndex++ {
		lines = updateNextGridRow(lines, rowIndex)
	}

	for rowIndex := 0; rowIndex < len(lines); rowIndex++ {
		splitters := findIndexesOfSplittersForRow(lines, rowIndex)
		for _, splitterIndex := range splitters {
			if string(lines[rowIndex-1][splitterIndex]) == "|" {
				total += 1
			}
		}
	}
	return total
}

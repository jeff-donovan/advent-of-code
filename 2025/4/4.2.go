package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Coords struct {
	i int
	j int
}

func parseInput(f *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func canRollOfPaperBeAccessed(grid []string, i int, j int) bool {
	numAdjacentRollsPaper := 0
	for _, i1 := range []int{i - 1, i, i + 1} {
		for _, j1 := range []int{j - 1, j, j + 1} {
			if (i1 != i || j1 != j) && isLegalPosition(grid, i1, j1) {
				if isRollOfPaper(grid, i1, j1) {
					numAdjacentRollsPaper += 1
				}
			}
		}
	}
	return numAdjacentRollsPaper < 4
}

func isLegalPosition(grid []string, i int, j int) bool {
	iLength := len(grid[0])
	jLength := len(grid)
	return 0 <= i && i < iLength && 0 <= j && j < jLength
}

func isRollOfPaper(grid []string, i int, j int) bool {
	return grid[j][i] == '@'
}

func makeNewGrid(grid []string, remove []Coords) []string {
	var newGrid []string
	for _, line := range grid {
		newGrid = append(newGrid, line)
	}

	for _, c := range remove {
		newGrid[c.j] = newGrid[c.j][0:c.i] + "x" + newGrid[c.j][c.i+1:]
	}
	return newGrid
}

func main() {
	// f, err := os.Open("C:/code/advent-of-code/2025/4/day_4_input.txt")
	f, err := os.Open("C:/code/advent-of-code/2025/4/day_4_test.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer f.Close()
	grid, err := parseInput(f)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	start := time.Now()

	total := 0
	var remove []Coords
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if isRollOfPaper(grid, i, j) && canRollOfPaperBeAccessed(grid, i, j) {
				// if isRollOfPaper(grid, i, j) {
				fmt.Printf("roll of paper! %d, %d\n", i, j)
				remove = append(remove, Coords{i, j})
				total += 1
			}
		}
	}

	newGrid := makeNewGrid(grid, remove)
	for _, line := range newGrid {
		fmt.Println(line)
	}

	fmt.Println("Answer: ", total)

	fmt.Println("took: ", time.Since(start))
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Coord struct {
	x int
	y int
}

func calculateArea(a, b Coord) int {
	answer := (int(math.Abs(float64(a.x-b.x))) + 1) * (int(math.Abs(float64(a.y-b.y))) + 1)
	fmt.Println("Area: ", answer)
	return answer
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

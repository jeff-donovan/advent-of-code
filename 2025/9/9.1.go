package main

import (
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func calculateArea(a, b Coord) int {
	if a.x == b.x || a.y == b.y {
		return -1
	}

	return int(math.Abs(float64(a.x-b.x))) * int(math.Abs(float64(a.y-b.y)))
}

func algorithm9_1(lines []string) int {
	total := 0

	var coords []Coord
	for _, l := range lines {
		numStrings := strings.Split(l, ",")
		x, _ := strconv.Atoi(numStrings[0])
		y, _ := strconv.Atoi(numStrings[1])
		coords = append(coords, Coord{x, y})
	}

	var areas []int
	for _, c1 := range coords {
		for _, c2 := range coords {
		}
	}

	return total
}

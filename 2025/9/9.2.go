package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func calculateMinX(coords []Coord) int {
	min := math.MaxInt
	for _, c := range coords {
		if c.x < min {
			min = c.x
		}
	}
	return min
}

func calculatemaxY(coords []Coord) int {
	max := 0
	for _, c := range coords {
		if c.y > max {
			max = c.y
		}
	}
	return max
}

func getAllAdjacentCoords(a, b Coord) []Coord {
	if a.x != b.x && a.y != b.y {
		panic("oops! assumption wrong, these are not matching")
	}

	if a.x == b.x {
		minY := a.y
		if b.y < minY {
			minY = b.y
		}

		maxY := a.y
		if b.y > maxY {
			maxY = b.y
		}

		var coords []Coord
		for y := minY; y <= maxY; y++ {
			coords = append(coords, Coord{a.x, y})
		}
		return coords
	}

	// otherwise a.y == b.y
	minX := a.x
	if b.x < minX {
		minX = b.x
	}

	maxX := a.x
	if b.x > maxX {
		maxX = b.x
	}

	var coords []Coord
	for x := minX; x <= maxX; x++ {
		coords = append(coords, Coord{x, a.y})
	}
	return coords
}

func makeAllCoordsMap(coords []Coord) map[Coord]struct{} {
	coordsMap := make(map[Coord]struct{})
	for i := 0; i < len(coords)-1; i++ {
		a := coords[i]
		b := coords[i+1]
		adjCoords := getAllAdjacentCoords(a, b)
		for _, c := range adjCoords {
			coordsMap[c] = struct{}{}
		}
	}
	return coordsMap
}

func isValidRectangle(coordsMap map[Coord]struct{}, a, b Coord) bool {
	return false
}

func algorithm9_2(lines []string) int {
	// plan
	//  - get all coords
	//  - make the coords map
	//  - figure out how to define a "valid" rectangle
	//  - loop through and only add the valid rectangle areas

	var coords []Coord
	for _, l := range lines {
		numStrings := strings.Split(l, ",")
		x, _ := strconv.Atoi(numStrings[0])
		y, _ := strconv.Atoi(numStrings[1])
		coords = append(coords, Coord{x, y})
	}

	coordsMap := makeAllCoordsMap(coords)

	var areas []int
	for _, c1 := range coords {
		for _, c2 := range coords {
			if isValidRectangle(coordsMap, c1, c2) {
				areas = append(areas, calculateArea(c1, c2))
			}
		}
	}

	fmt.Println("areas before: ", areas)

	max := 0
	for _, a := range areas {
		if a > max {
			max = a
		}
	}

	// fmt.Println("areas after: ", areas)

	return max
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculateMaxX(coords []Coord) int {
	max := 0
	for _, c := range coords {
		if c.x > max {
			max = c.x
		}
	}
	return max
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

func makeCoordsGrid(coords []Coord) []string {
	maxX := calculateMaxX(coords)
	maxY := calculatemaxY(coords)

	var grid []string
	for y := 0; y <= maxY; y++ {
		grid = append(grid, strings.Repeat(".", maxX+1))
	}
	return grid
}

func drawCoords(grid []string, coords []Coord) []string {
	var newGrid []string
	newGrid = append(newGrid, grid...)

	for _, c := range coords {
		newRow := newGrid[c.y]
		newRow = newRow[:c.x] + "#" + newRow[c.x+1:]
		newGrid[c.y] = newRow
	}

	return newGrid
}

func drawAllCoords(grid []string, coords []Coord) []string {
	newGrid := drawCoords(grid, coords)

	// draw all horizontal
	for y, row := range newGrid {
		firstPound := strings.Index(row, "#")
		if firstPound == -1 {
			continue
		}

		lastPound := strings.LastIndex(row, "#")
		newGrid[y] = row[0:firstPound] + strings.Repeat("#", (lastPound-firstPound)+1) + row[lastPound+1:]
	}

	// draw all vertical
	for x := 0; x < len(newGrid[0]); x++ {
		col := ""
		for y := 0; y < len(newGrid); y++ {
			col += string(newGrid[y][x])
		}

		firstPound := strings.Index(col, "#")
		if firstPound == -1 {
			continue
		}

		lastPound := strings.LastIndex(col, "#")
		for y := firstPound; y <= lastPound; y++ {
			newGrid[y] = newGrid[y][0:x] + "#" + newGrid[y][x+1:]
		}
	}

	return newGrid
}

func getAllRectangleCoords(a, b Coord) []Coord {
	minX := a.x
	if b.x < minX {
		minX = b.x
	}

	maxX := a.x
	if b.x > maxX {
		maxX = b.x
	}

	minY := a.y
	if b.y < minY {
		minY = b.y
	}

	maxY := a.y
	if b.y > maxY {
		maxY = b.y
	}

	var coords []Coord
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			coords = append(coords, Coord{x, y})
		}
	}
	return coords
}

func isValidRectangle(grid []string, a, b Coord) bool {
	rectangleCoords := getAllRectangleCoords(a, b)
	for _, c := range rectangleCoords {
		exists := string(grid[c.y][c.x]) == "#"
		if !exists {
			return false
		}
	}
	return true
}

func algorithm9_2(lines []string) int {
	// plan
	//  - get all coords
	//  - draw the coords grid
	//  - figure out how to define a "valid" rectangle
	//  - loop through and only add the valid rectangle areas

	var coords []Coord
	for _, l := range lines {
		numStrings := strings.Split(l, ",")
		x, _ := strconv.Atoi(numStrings[0])
		y, _ := strconv.Atoi(numStrings[1])
		coords = append(coords, Coord{x, y})
	}

	// coordsMap := makeAllCoordsMap(coords)
	// for c, _ := range coordsMap {
	// 	fmt.Println(c)
	// }
	grid := makeCoordsGrid(coords)
	grid = drawAllCoords(grid, coords)
	for _, g := range grid {
		fmt.Println(g)
	}

	var areas []int
	for _, c1 := range coords {
		for _, c2 := range coords {
			if isValidRectangle(grid, c1, c2) {
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

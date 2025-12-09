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

func drawCoordsAndEdges(grid []string, coords []Coord) []string {
	var newGrid []string
	newGrid = append(newGrid, grid...)

	for i := 0; i < len(coords); i++ {
		a := coords[i]

		var b Coord
		if i+1 == len(coords) {
			b = coords[0]
		} else {
			b = coords[i+1]
		}

		if isHorizontal(a, b) {
			newRow := newGrid[a.y]
			minX := min(a.x, b.x)
			maxX := max(a.x, b.x)
			newRow = newRow[:minX] + strings.Repeat("#", (maxX-minX)+1) + newRow[maxX+1:]
			newGrid[a.y] = newRow
		}

		if isVertical(a, b) {
			minY := min(a.y, b.y)
			maxY := max(a.y, b.y)
			for y := minY; y <= maxY; y++ {
				newRow := newGrid[y]
				newRow = newRow[:a.x] + "#" + newRow[a.x+1:]
				newGrid[y] = newRow
			}
		}
	}

	return newGrid
}

func isHorizontal(a, b Coord) bool {
	return a.y == b.y
}

func isVertical(a, b Coord) bool {
	return a.x == b.x
}

func makeHorizontalRanges(gridWithCoords []string) [][]Coord {
	var ranges [][]Coord
	// draw all horizontal
	for y, row := range gridWithCoords {
		firstPound := strings.Index(row, "#")
		if firstPound == -1 {
			ranges = append(ranges, nil)
			continue
		}

		lastPound := strings.LastIndex(row, "#")
		ranges = append(ranges, []Coord{{firstPound, y}, {lastPound, y}})
	}
	return ranges
}

func makeVerticalRanges(gridWithCoords []string) [][]Coord {
	var ranges [][]Coord
	// draw all vertical
	for x := 0; x < len(gridWithCoords[0]); x++ {
		firstPound := -1
		for y := 0; y < len(gridWithCoords); y++ {
			if gridWithCoords[y][x] == '#' {
				firstPound = y
				break
			}
		}

		if firstPound == -1 {
			ranges = append(ranges, nil)
			continue
		}

		lastPound := -1
		for y := len(gridWithCoords) - 1; y >= firstPound; y-- {
			if gridWithCoords[y][x] == '#' {
				lastPound = y
				break
			}
		}

		ranges = append(ranges, []Coord{{x, firstPound}, {x, lastPound}})
	}
	return ranges
}

func areHorizontalsInRange(horizontalRanges [][]Coord, a, b Coord) bool {
	// top := []Coord{{a.x, a.y}, {b.x, a.y}}
	ayRange := horizontalRanges[a.y]
	isTopInRange := ayRange[0].x <= a.x && a.x <= ayRange[1].x && ayRange[0].x <= b.x && b.x <= ayRange[1].x
	if !isTopInRange {
		return false
	}

	// bottom := []Coord{{a.x, b.y}, {b.x, b.y}}
	byRange := horizontalRanges[b.y]
	isBottomInRange := byRange[0].x <= a.x && a.x <= byRange[1].x && byRange[0].x <= b.x && b.x <= byRange[1].x
	return isBottomInRange
}

func areVerticalsInRange(verticalRanges [][]Coord, a, b Coord) bool {
	// left := []Coord{{a.x, a.y}, {a.x, b.y}}
	axRange := verticalRanges[a.x]
	isLeftInRange := axRange[0].y <= a.y && a.y <= axRange[1].y && axRange[0].y <= b.y && b.y <= axRange[1].y
	if !isLeftInRange {
		return false
	}

	// right := []Coord{{b.x, a.y}, {b.x, b.y}}
	bxRange := verticalRanges[b.x]
	isRightInRange := bxRange[0].y <= a.y && a.y <= bxRange[1].y && bxRange[0].y <= b.y && b.y <= bxRange[1].y
	return isRightInRange
}

// func drawAllCoords(grid []string, coords []Coord) []string {
// 	newGrid := drawCoords(grid, coords)

// 	// draw all horizontal
// 	for y, row := range newGrid {
// 		firstPound := strings.Index(row, "#")
// 		if firstPound == -1 {
// 			continue
// 		}

// 		lastPound := strings.LastIndex(row, "#")
// 		newGrid[y] = row[0:firstPound] + strings.Repeat("#", (lastPound-firstPound)+1) + row[lastPound+1:]
// 	}

// 	// draw all vertical
// 	for x := 0; x < len(newGrid[0]); x++ {
// 		col := ""
// 		for y := 0; y < len(newGrid); y++ {
// 			col += string(newGrid[y][x])
// 		}

// 		firstPound := strings.Index(col, "#")
// 		if firstPound == -1 {
// 			continue
// 		}

// 		lastPound := strings.LastIndex(col, "#")
// 		for y := firstPound; y <= lastPound; y++ {
// 			newGrid[y] = newGrid[y][0:x] + "#" + newGrid[y][x+1:]
// 		}
// 	}

// 	return newGrid
// }

func getRectangleHorizontalRanges(a, b Coord) [][]Coord {
	top := []Coord{{a.x, a.y}, {b.x, a.y}}
	bottom := []Coord{{a.x, b.y}, {b.x, b.y}}
	return [][]Coord{top, bottom}
}

func getRectangleVerticalRanges(a, b Coord) [][]Coord {
	left := []Coord{{a.x, a.y}, {a.x, b.y}}
	right := []Coord{{b.x, a.y}, {b.x, b.y}}
	return [][]Coord{left, right}
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

func isValidRectangle(horizontalRanges [][]Coord, verticalRanges [][]Coord, a, b Coord) bool {
	return areHorizontalsInRange(horizontalRanges, a, b) && areVerticalsInRange(verticalRanges, a, b)
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
	fmt.Println("finished making coords")

	grid := makeCoordsGrid(coords)
	fmt.Println("finished making grid")

	grid = drawCoordsAndEdges(grid, coords)
	fmt.Println("finished drawing coords AND edges on grid")

	horizontalRanges := makeHorizontalRanges(grid)
	fmt.Println("finished making horizontal ranges")

	verticalRanges := makeVerticalRanges(grid)
	fmt.Println("finished making vertical ranges")

	// // now we need to add in the outer Xs so that we fill in the remaining ranges
	// grid = drawRanges(grid, horizontalRanges, verticalRanges)
	// horizontalRanges = makeHorizontalRanges(grid)
	// verticalRanges = makeVerticalRanges(grid)

	max := 0
	for _, c1 := range coords {
		for _, c2 := range coords {
			if isValidRectangle(horizontalRanges, verticalRanges, c1, c2) {
				area := calculateArea(c1, c2)
				if area > max {
					max = area
				}
			}
		}
	}

	// max := 0
	// for _, a := range areas {
	// 	if a > max {
	// 		max = a
	// 	}
	// }

	return max
	// return 0
}

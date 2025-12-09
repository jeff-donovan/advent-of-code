package main

import (
	"fmt"
	"strconv"
	"strings"
)

func algorithm9_1(lines []string) int {
	var coords []Coord
	for _, l := range lines {
		numStrings := strings.Split(l, ",")
		x, _ := strconv.Atoi(numStrings[0])
		y, _ := strconv.Atoi(numStrings[1])
		fmt.Println("x: ", x)
		fmt.Println("y: ", y)
		coords = append(coords, Coord{x, y})
	}

	for _, c := range coords {
		fmt.Println(c)
	}

	var areas []int
	for _, c1 := range coords {
		for _, c2 := range coords {
			areas = append(areas, calculateArea(c1, c2))
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

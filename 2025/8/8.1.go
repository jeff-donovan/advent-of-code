package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type PairWithDistance struct {
	a        JunctionBox
	b        JunctionBox
	distance float64
}

type Pair struct {
	a JunctionBox
	b JunctionBox
}

type JunctionBox struct {
	x int
	y int
	z int
}

func makeJunctionBoxes(lines []string) []JunctionBox {
	var result []JunctionBox

	for _, l := range lines {
		coords := strings.Split(l, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		result = append(result, JunctionBox{x, y, z})
	}

	return result
}

func distanceBetweenBoxes(a, b JunctionBox) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y) + (a.z-b.z)*(a.z-b.z)))
}

func algorithm8_1(lines []string) int {
	total := 0

	// plan
	//  - parse the input and get a list of junction boxes
	//  - make all possible pairs of junction boxes (except for duplicates since the distance is 0)
	//  - calculate the distance between each pair
	//  - sort the pairs by distance in ascending order

	junctionBoxes := makeJunctionBoxes(lines)
	pairs := make(map[Pair]float64)
	for _, a := range junctionBoxes {
		for _, b := range junctionBoxes {
			if a == b {
				continue
			}

			abPair := Pair{a, b}
			_, abExists := pairs[abPair]
			if abExists {
				continue
			}

			baPair := Pair{b, a}
			_, baExists := pairs[baPair]
			if baExists {
				continue
			}

			pairs[Pair{a, b}] = distanceBetweenBoxes(a, b)
		}
	}

	var pairsWithDistances []PairWithDistance
	for pair, distance := range pairs {
		pairsWithDistances = append(pairsWithDistances, PairWithDistance{pair.a, pair.b, distance})
	}
	sort.Slice(pairsWithDistances, func(i, j int) bool {
		return pairsWithDistances[i].distance < pairsWithDistances[j].distance
	})

	for _, p := range pairsWithDistances {
		fmt.Printf("%f : (%v, %v)\n", p.distance, p.a, p.b)
	}

	return total
}

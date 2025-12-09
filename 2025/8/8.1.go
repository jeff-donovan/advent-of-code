package main

import (
	"fmt"
	"math"
	"slices"
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

func makePairsWithDistances(junctionBoxes []JunctionBox) []PairWithDistance {
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

	return pairsWithDistances
}

func algorithm8_1(lines []string, numConnections int) int {
	// plan
	//  - parse the input and get a list of junction boxes
	//  - make all possible pairs of junction boxes (except for duplicates since the distance is 0)
	//  - calculate the distance between each pair
	//  - sort the pairs by distance in ascending order

	junctionBoxes := makeJunctionBoxes(lines)
	pairs := makePairsWithDistances(junctionBoxes)

	// plan
	//  - number of connections
	//  - make list of circuits of JunctionBoxes
	//  - consider if we need to merge circuits together? does that happen?

	connectionsLeft := numConnections
	var circuits [][]JunctionBox

	for _, p := range pairs {
		if connectionsLeft == 0 {
			break
		}

		connectionsLeft--

		circuitIndexA := -1
		circuitIndexB := -1
		for i, c := range circuits {
			for _, jb := range c {
				if p.a == jb {
					circuitIndexA = i
				}
				if p.b == jb {
					circuitIndexB = i
				}
			}
		}

		// both junction boxes are in a circuit, skip IF we can't combine
		if circuitIndexA != -1 && circuitIndexB != -1 {
			if circuitIndexA != circuitIndexB {
				fmt.Println("connections left: ", connectionsLeft)
				fmt.Println("Circuits before: ")
				for _, c := range circuits {
					fmt.Println(c)
				}
				fmt.Println()

				var newCircuits [][]JunctionBox
				for i, c := range circuits {
					if i != circuitIndexA && i != circuitIndexB {
						newCircuits = append(newCircuits, c)
					}
				}
				newCircuits = append(newCircuits, slices.Concat(circuits[circuitIndexA], circuits[circuitIndexB]))
				circuits = newCircuits

				fmt.Println("Circuits after: ")
				for _, c := range circuits {
					fmt.Println(c)
				}
				fmt.Println()
			}
			continue
		}

		// neither junction box is in a circuit, make a new one
		if circuitIndexA == -1 && circuitIndexB == -1 {
			circuits = append(circuits, []JunctionBox{p.a, p.b})
		}

		if circuitIndexA != -1 {
			circuits[circuitIndexA] = append(circuits[circuitIndexA], p.b)
		}

		if circuitIndexB != -1 {
			circuits[circuitIndexB] = append(circuits[circuitIndexB], p.a)
		}
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	fmt.Println("Circuit #1: ", circuits[0])
	fmt.Println("Circuit #2: ", circuits[1])
	fmt.Println("Circuit #3: ", circuits[2])

	for _, c := range circuits {
		fmt.Println(len(c), c)
	}

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

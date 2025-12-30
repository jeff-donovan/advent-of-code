package main

import "fmt"

type Clicks struct {
	currentReq JoltageRequirement
	n          int
}

func bfs(machine *Machine) int {
	var initialJoltage JoltageRequirement
	for range machine.requirements {
		initialJoltage = append(initialJoltage, 0)
	}

	queue := []Clicks{{initialJoltage, 0}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if areEqual(machine.requirements, current.currentReq) {
			return current.n
		}

		if isImpossible(machine, current.currentReq) {
			continue
		}

		for _, b := range machine.buttons {
			nextReq := makeNextJoltageRequirement(current.currentReq, b)
			if !isImpossible(machine, nextReq) {
				queue = append(queue, Clicks{nextReq, current.n + 1})
			}
		}
	}

	return 0
}

func algorithm10_2_3(lines []string) int {
	total := 0
	machines := makeMachines(lines)
	for i, m := range machines {
		val := bfs(&m)
		fmt.Printf("calculated min button clicks for machine %d: %d\n", i, val)
		total += val
		// fmt.Println(m)
	}

	return total
}

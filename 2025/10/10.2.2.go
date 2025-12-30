package main

import (
	"fmt"
	"math"
)

func isImpossible(machine *Machine, current JoltageRequirement) bool {
	for i := range machine.requirements {
		if current[i] > machine.requirements[i] {
			return true
		}
	}

	return false
}

func areEqual(a, b JoltageRequirement) bool {
	if len(a) != len(b) {
		fmt.Println("unexpectedly different lengths!")
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func dfsMinimumClicks(machine Machine) int {
	var initialJoltage JoltageRequirement
	for range machine.requirements {
		initialJoltage = append(initialJoltage, 0)
	}
	return dfs(&machine, 0, initialJoltage)
}

func dfs(machine *Machine, n int, currentJoltage JoltageRequirement) int {
	if areEqual(machine.requirements, currentJoltage) {
		return n
	}

	if n >= machine.minVal {
		return n
	}

	if isImpossible(machine, currentJoltage) {
		return math.MaxInt
	}

	minClicks := math.MaxInt
	for _, b := range machine.buttons {
		nextReq := makeNextJoltageRequirement(currentJoltage, b)
		minClicks = min(minClicks, dfs(machine, n+1, nextReq))
	}

	machine.minVal = min(machine.minVal, minClicks)

	return min(machine.minVal, minClicks)
}

func algorithm10_2_2(lines []string) int {
	total := 0
	machines := makeMachines(lines)
	for i, m := range machines {
		val := dfsMinimumClicks(m)
		fmt.Printf("calculated min button clicks for machine %d: %d\n", i, val)
		total += val
		// fmt.Println(m)
	}

	return total
}

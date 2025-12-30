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
	key := makeJoltageReqKey(currentJoltage)

	minVal, exists := machine.minValsMap[key]
	if exists {
		if minVal == math.MaxInt {
			return math.MaxInt
		}
		return n + minVal
	}

	if areEqual(machine.requirements, currentJoltage) {
		return n
	}

	if isImpossible(machine, currentJoltage) {
		return math.MaxInt
	}

	minClicksFromHere := math.MaxInt
	for _, b := range machine.buttons {
		nextReq := makeNextJoltageRequirement(currentJoltage, b)
		minClicksFromHere = min(minClicksFromHere, dfs(machine, 1, nextReq))
	}

	machine.minValsMap[key] = minClicksFromHere

	if minClicksFromHere == math.MaxInt {
		return math.MaxInt
	}
	return n + minClicksFromHere
}

func algorithm10_2_2(lines []string) int {
	total := 0
	machines := makeMachines(lines)
	for i, m := range machines {
		val := dfsMinimumClicks(m)
		fmt.Printf("calculated min button clicks for machine %d: %d | machine minVal: %d\n", i, val, m.minVal)
		total += val
		// fmt.Println(m)
	}

	return total
}

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func calculateFewestButtonClicksForJoltageRequirements(n int, machine *Machine, currentJoltage JoltageRequirement) int {
	// fmt.Println("map: ", machine.isImpossiblePathMap)
	if areRequirementsEqual(machine.requirements, currentJoltage) {
		return n
	}

	if n > machine.minVal {
		return n
	}

	if isImpossiblePath(n, machine, currentJoltage) {
		return math.MaxInt
	}

	var stack []JoltageRequirement
	for _, b := range machine.buttons {
		nextReq := makeNextJoltageRequirement(currentJoltage, b)
		stack = append(stack, nextReq)
	}

	minVal := math.MaxInt
	for _, r := range stack {
		val := calculateFewestButtonClicksForJoltageRequirements(n+1, machine, r)
		if val < minVal {
			minVal = val
		}
	}

	if minVal == math.MaxInt || minVal > machine.minVal {
		addJoltageRequirementToImpossibleMap(n, machine, currentJoltage)
		return minVal
	}

	machine.minVal = minVal
	return minVal
}

func calculateFewestButtonClicksRemaining(machine *Machine, current JoltageRequirement) int {
	// fmt.Println("current: ", current)
	if areRequirementsEqual(machine.requirements, current) {
		return 0
	}

	if isImpossiblePath(0, machine, current) {
		return math.MaxInt
	}

	// fmt.Println("NOT! impossible path!")
	var produces []int
	// TODO: trim Infinity
	for _, click := range machine.buttons {
		val := calculateFewestButtonClicksRemaining(machine, makeNextJoltageRequirement(current, click))
		if val != math.MaxInt {
			// fmt.Println("val is: ", val)
			produces = append(produces, 1+val)
			fmt.Println("length of produces is: ", len(produces))
		}
	}

	minVal := math.MaxInt
	for _, v := range produces {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func makeNextJoltageRequirement(current JoltageRequirement, click Button) JoltageRequirement {
	var next JoltageRequirement
	next = append(next, current...)

	for _, i := range click {
		next[i]++
	}

	return next
}

func makeEndResultJoltageRequirements(machine Machine, clicks []Button) JoltageRequirement {
	// TODO: add memoization
	generatedRequirements := make(JoltageRequirement, len(machine.requirements))
	// fmt.Println("generatedRequirements before: ", generatedRequirements)
	// fmt.Println("clicks: ", clicks)
	for _, click := range clicks {
		for _, i := range click {
			generatedRequirements[i]++
		}
	}
	// fmt.Println("generatedRequirements after: ", generatedRequirements)
	return generatedRequirements
}

func addJoltageRequirementToImpossibleMap(n int, machine *Machine, current JoltageRequirement) {
	key := makeJoltageReqKey(current)
	_, exists := machine.isImpossiblePathMap[key]
	if !exists {
		machine.isImpossiblePathMap[key] = make(map[int]struct{})
	}

	machine.isImpossiblePathMap[key][n] = struct{}{}
}

func isImpossiblePath(n int, machine *Machine, current JoltageRequirement) bool {
	key := makeJoltageReqKey(current)
	nMap, exists := machine.isImpossiblePathMap[key]
	if exists {
		_, existsForN := nMap[n]
		if existsForN {
			return true
		}
	}

	if len(machine.requirements) != len(current) {
		panic("unexpectedly different lengths!")
	}

	for i := range machine.requirements {
		if current[i] > machine.requirements[i] {
			addJoltageRequirementToImpossibleMap(n, machine, current)
			return true
		}
	}

	return false
}

func makeJoltageReqKey(req JoltageRequirement) string {
	var reqsAsStrings []string
	for _, val := range req {
		reqsAsStrings = append(reqsAsStrings, strconv.Itoa(val))
	}
	return strings.Join(reqsAsStrings, ",")
}

func areRequirementsEqual(a, b JoltageRequirement) bool {
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

func algorithm10_2(lines []string) int {
	total := 0

	machines := makeMachines(lines)
	for i, m := range machines {
		var initialJoltage JoltageRequirement
		for range m.requirements {
			initialJoltage = append(initialJoltage, 0)
		}

		val := calculateFewestButtonClicksForJoltageRequirements(0, &m, initialJoltage)
		fmt.Printf("calculated min button clicks for machine %d: %d\n", i, val)
		total += val
		// fmt.Println(m)
	}
	// fmt.Println("min: ", calculateFewestButtonClicksForJoltageRequirements(machines[0]))
	// fmt.Println("equal? ", areRequirementsEqual(machines[0].requirements, JoltageRequirement{3, 5, 4, 7}))

	// m := machines[0]
	// current := JoltageRequirement{3, 0, 0, 0}
	// fmt.Println(makeNextJoltageRequirement(current, Button{3}))

	// plan
	//  - once a value is greater than the corresponding value in the requirements, we've found an impossible permutation path
	//  - need to use recursion or maybe a stack

	// machine := machines[0]
	// result := calculateFewestButtonClicks(machine)
	// fmt.Println("machine 0: ", result)

	return total
}

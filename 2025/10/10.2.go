package main

import (
	"fmt"
	"math"
)

func calculateFewestButtonClicksForJoltageRequirements(machine Machine) int {
	// set up initial joltagereq
	// call calculateFewestButtonClicksRemaining

	var initialJoltage JoltageRequirement
	for range machine.requirements {
		initialJoltage = append(initialJoltage, 0)
	}

	stack := []JoltageRequirement{initialJoltage}
	n := 1
	for {
		var reqsToCheck []JoltageRequirement
		reqsToCheck = append(reqsToCheck, stack...)
		stack = nil

		for _, r := range reqsToCheck {
			for _, b := range machine.buttons {
				nextReq := makeNextJoltageRequirement(r, b)
				if areRequirementsEqual(machine.requirements, nextReq) {
					return n
				}

				if isImpossiblePath(machine.requirements, nextReq) {
					continue
				}

				stack = append(stack, nextReq)
			}
		}

		n++
	}
}

func calculateFewestButtonClicksRemaining(machine Machine, current JoltageRequirement) int {
	// fmt.Println("current: ", current)
	if areRequirementsEqual(machine.requirements, current) {
		return 0
	}

	if isImpossiblePath(machine.requirements, current) {
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

func isImpossiblePath(req JoltageRequirement, current JoltageRequirement) bool {
	if len(req) != len(current) {
		panic("unexpectedly different lengths!")
	}

	for i := range req {
		if current[i] > req[i] {
			return true
		}
	}

	return false
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
	for _, m := range machines {
		total += calculateFewestButtonClicksForJoltageRequirements(m)
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

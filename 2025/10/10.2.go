package main

import "fmt"

func calculateFewestButtonClicksForJoltageRequirements(machine Machine) int {
	n := 1
	for {
		permutations := generatePermutations(machine.buttons, n)
		for _, p := range permutations {
			result := makeEndResultJoltageRequirements(machine, p)
			if areRequirementsEqual(machine.requirements, result) {
				return n
			}
		}
		n++
	}
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
	}

	// machine := machines[0]
	// result := calculateFewestButtonClicks(machine)
	// fmt.Println("machine 0: ", result)

	return total
}

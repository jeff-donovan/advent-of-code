package main

import "fmt"

func algorithm10_2(lines []string) int {
	total := 0

	machines := makeMachines(lines)
	for _, m := range machines {
		fmt.Println(m)
	}

	// machine := machines[0]
	// result := calculateFewestButtonClicks(machine)
	// fmt.Println("machine 0: ", result)

	return total
}

package main

func algorithm10_1(lines []string) int {
	total := 0

	machines := makeMachines(lines)
	for _, m := range machines {
		total += calculateFewestButtonClicks(m)
	}

	// machine := machines[0]
	// result := calculateFewestButtonClicks(machine)
	// fmt.Println("machine 0: ", result)

	return total
}

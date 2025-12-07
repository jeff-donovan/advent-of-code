package main

func algorithm6_2(lines []string) int {
	total := 0
	for _, p := range mapInputToMathProblems(cleanLines(lines)) {
		if p.operation == "*" {
			total += solveMultiplication(p)
		}
		if p.operation == "+" {
			total += solveAddition(p)
		}
	}
	return total
}

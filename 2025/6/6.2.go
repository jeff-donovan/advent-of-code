package main

import (
	"strconv"
)

func mapInputToMathProblemsPart2(lines []string) []MathProblem {
	var problems []MathProblem
	operationIndex := len(lines) - 1

	var nums []int
	for n := len(lines[0]) - 1; n >= 0; n-- {
		numString := ""
		for j, l := range lines {
			cleanedChar := string(l[n])
			if cleanedChar == " " {
				cleanedChar = ""
			}

			if j != operationIndex {
				numString += cleanedChar
				continue
			}

			// we're now at the operation line so let's turn the string into an int
			// but only if it's not an empty string, otherwise that gets cast as 0
			if numString != "" {
				num, _ := strconv.Atoi(numString)
				nums = append(nums, num)
				numString = ""
			}

			// base case - we've found an operation!
			if cleanedChar != "" {
				var pNums []int
				pNums = append(pNums, nums...)
				problems = append(problems, MathProblem{cleanedChar, pNums})
				nums = nil
				break
			}
		}
	}

	return problems
}

func algorithm6_2(lines []string) int {
	total := 0

	for _, p := range mapInputToMathProblemsPart2(lines) {
		if p.operation == "*" {
			total += solveMultiplication(p)
		}
		if p.operation == "+" {
			total += solveAddition(p)
		}
	}
	return total
}

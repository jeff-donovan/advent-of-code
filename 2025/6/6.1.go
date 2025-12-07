package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MathProblem struct {
	operation string
	nums      []int
}

func mapInputToMathProblems(lines [][]string) []MathProblem {
	var problems []MathProblem
	operationIndex := len(lines) - 1
	numProblems := len(lines[0])
	for n := 0; n < numProblems; n++ {
		operation := lines[operationIndex][n]
		var nums []int
		for j := 0; j < operationIndex; j++ {
			num, _ := strconv.Atoi(lines[j][n])
			nums = append(nums, num)
		}
		problems = append(problems, MathProblem{operation, nums})
	}
	return problems
}

func cleanLines(lines []string) [][]string {
	var result [][]string
	for _, l := range lines {
		var final []string

		// have to clean otherwise we're left with empty strings in the lists, in addition to the number strings
		for _, s := range strings.Split(l, " ") {
			if s != "" {
				final = append(final, s)
			}
		}
		result = append(result, final)
	}

	return result
}

func algorithm6_1(lines []string) int {
	problems := mapInputToMathProblems(cleanLines(lines))
	for _, p := range problems {
		fmt.Println(p)
	}
	return 0
}

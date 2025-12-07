package main

import (
	"fmt"
	"strings"
)

type MathProblem struct {
	operation string
	nums      []int
}

func mapInputToMathProblems(lines []string) []MathProblem {
	return nil
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
	problems := cleanLines(lines)
	// problems := mapInputToMathProblems(lines)
	for _, p := range problems {
		fmt.Println(p)
	}
	return 0
}

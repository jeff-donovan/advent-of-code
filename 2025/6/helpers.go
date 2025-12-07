package main

import (
	"bufio"
	"os"
)

type MathProblem struct {
	operation string
	nums      []int
}

func solveAddition(problem MathProblem) int {
	total := 0
	for _, n := range problem.nums {
		total += n
	}
	return total
}

func solveMultiplication(problem MathProblem) int {
	total := 1
	for _, n := range problem.nums {
		total *= n
	}
	return total
}

func parseInput(f *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

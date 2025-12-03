package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

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

func calculateLargestPossibleBankVoltage(bank string) int {
	for tens := 9; tens > 0; tens-- {
		for i, char := range bank {
			digit, _ := strconv.Atoi(string(char))
			if digit == tens && i < len(bank)-1 {
				return 10*tens + getMaxDigitInString(bank[i+1:])
			}
		}
	}
	return 0
}

func getMaxDigitInString(s string) int {
	current := 0
	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		if digit > current {
			current = digit
		}
	}
	return current
}

func main() {
	f, err := os.Open("C:/code/advent-of-code/2025/3/day_3_input.txt")
	// f, err := os.Open("C:/code/advent-of-code/2025/3/day_3_test.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer f.Close()
	lines, err := parseInput(f)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	start := time.Now()

	total := 0
	for _, bank := range lines {
		total += calculateLargestPossibleBankVoltage(bank)
	}

	fmt.Println("Answer: ", total)

	fmt.Println("took: ", time.Since(start))
}

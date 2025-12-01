package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	// f, err := os.Open("C:/code/advent-of-code/2025/1/day_1_input.txt")
	f, err := os.Open("C:/code/advent-of-code/2025/1/day_1_test.txt")
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

	for _, code := range lines {
		fmt.Println(code)
	}

	fmt.Println("took: ", time.Since(start))
}

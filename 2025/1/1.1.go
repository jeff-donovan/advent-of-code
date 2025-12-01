package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func calculatePassword(lines []string) int {
	current := 50
	password := 0

	for _, line := range lines {
		num, _ := strconv.Atoi(line[1:])
		if strings.HasPrefix(line, "R") {
			current = (current + num) % 100
		} else {
			current = (current - num) % 100
		}

		if current == 0 {
			password += 1
		}
	}

	return password
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

	fmt.Println(calculatePassword(lines))

	fmt.Println("took: ", time.Since(start))
}

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
			diff := 100 - current
			if num > diff {
				password += 1
			}
			additionalClicks := (num - diff) / 100
			current = (current + num) % 100
			if current == 0 {
				additionalClicks = additionalClicks - 1
			}
			password += additionalClicks
		} else {
			diff := current
			if num > diff {
				password += 1
			}
			additionalClicks := abs((num - diff) / 100)
			current = (current - num) % 100
			if current == 0 {
				additionalClicks = additionalClicks - 1
			}
			password += additionalClicks
		}

		if current == 0 {
			password += 1
		}
	}

	return password
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

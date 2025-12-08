package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Open("C:/code/advent-of-code/2025/8/day_8_input.txt")
	// f, err := os.Open("C:/code/advent-of-code/2025/8/day_8_test.txt")
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

	numConnections := 1000 // TODO: change for test vs. input
	total := algorithm8_1(lines, numConnections)

	fmt.Println("Answer: ", total)

	fmt.Println("took: ", time.Since(start))
}

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// f, err := os.Open("C:/code/advent-of-code/2025/7/day_7_input.txt")
	f, err := os.Open("C:/code/advent-of-code/2025/7/day_7_test.txt")
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

	total := algorithm7_2(lines)

	fmt.Println("Answer: ", total)

	fmt.Println("took: ", time.Since(start))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Range struct {
	Start int
	End   int
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

func parseRanges(lines []string) []Range {
	input := lines[0]
	rangeStrings := strings.Split(input, ",")
	var ranges []Range
	for _, rs := range rangeStrings {
		startEnd := strings.Split(rs, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])
		ranges = append(ranges, Range{Start: start, End: end})
	}

	return ranges
}

func main() {
	// f, err := os.Open("C:/code/advent-of-code/2025/2/day_2_input.txt")
	f, err := os.Open("C:/code/advent-of-code/2025/2/day_2_test.txt")
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

	ranges := parseRanges(lines)

	for i, r := range ranges {
		fmt.Println(i)
		fmt.Println(r)
		fmt.Println()
	}

	fmt.Println("took: ", time.Since(start))
}

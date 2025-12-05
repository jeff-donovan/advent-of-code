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
	start int
	end   int
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

func getRanges(lines []string) []Range {
	var ranges []Range
	for _, line := range lines {
		if strings.Contains(line, "-") {
			startEndStrings := strings.Split(line, "-")
			start, _ := strconv.Atoi(startEndStrings[0])
			end, _ := strconv.Atoi(startEndStrings[1])
			ranges = append(ranges, Range{start, end})
		}
	}
	return ranges
}

func getAvailableIngredientIds(lines []string) []int {
	var ids []int
	for _, line := range lines {
		if !strings.Contains(line, "-") {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}
	return ids
}

func main() {
	f, err := os.Open("C:/code/advent-of-code/2025/5/day_5_input.txt")
	// f, err := os.Open("C:/code/advent-of-code/2025/5/day_5_test.txt")
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

	ranges := getRanges(lines)
	for _, r := range ranges {
		fmt.Println(r)
	}
	ids := getAvailableIngredientIds(lines)

	total := 0

	for _, id := range ids {
		for _, r := range ranges {
			if r.start <= id && id <= r.end {
				total += 1
				break
			}
		}
	}

	fmt.Println("Answer: ", total)

	fmt.Println("took: ", time.Since(start))
}

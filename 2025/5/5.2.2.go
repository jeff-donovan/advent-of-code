package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func sortRanges(ranges []Range) {
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.start - b.start
	})
}

func isOverlapping(a, b Range) bool {
	return a.start <= b.start && b.start <= a.end
}

func mergeRanges(a, b Range) Range {
	minStart := a.start
	if b.start < minStart {
		minStart = b.start
	}

	maxEnd := a.end
	if b.end > maxEnd {
		maxEnd = b.end
	}

	return Range{minStart, maxEnd}
}

func mergeRangesUntilOne(ranges []Range) []Range {
	if len(ranges) <= 1 {
		return ranges
	}

	newFirstMergedRange := mergeRanges(ranges[0], ranges[1])
	if len(ranges) == 2 {
		return []Range{newFirstMergedRange}
	}

	return mergeRangesUntilOne(slices.Concat([]Range{newFirstMergedRange}, ranges[2:]))
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
	sortRanges(ranges)

	var overlappingRanges []Range

	for _, r1 := range ranges {
		for _, r2 := range ranges {
			if isOverlapping(r1, r2) {
				overlappingRanges = append(overlappingRanges, mergeRanges(r1, r2))
			}
		}
	}

	var newRanges []Range
	finalRangesMap := make(map[int]struct{})
	for _, r1 := range overlappingRanges {
		_, exists := finalRangesMap[r1.start]
		if exists {
			continue
		}
		finalRangesMap[r1.start] = struct{}{}

		var rangesWithSameStart []Range
		for _, r2 := range overlappingRanges {
			if r2.start == r1.start {
				rangesWithSameStart = append(rangesWithSameStart, r2)
			}
		}

		newRanges = slices.Concat(newRanges, mergeRangesUntilOne(rangesWithSameStart))
	}

	for _, r := range newRanges {
		fmt.Println(r)
	}

	fmt.Println("Num Ranges: ", len(ranges))
	fmt.Println("Num new ranges: ", len(newRanges))

	// fmt.Println("Answer: ", len(freshIds))

	fmt.Println("took: ", time.Since(start))
}

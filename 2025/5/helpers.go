package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
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

func isOverlapping(a, b Range) bool {
	return (a.start <= b.start && b.start <= a.end) || (b.start <= a.start && a.start <= b.end)
}

func dedupe(ranges []Range) []Range {
	var finalRanges []Range

	for _, r1 := range ranges {
		seen := false
		for _, r2 := range finalRanges {
			if r1.start == r2.start && r1.end == r2.end {
				seen = true
				break
			}
		}
		if !seen {
			finalRanges = append(finalRanges, r1)
		}
	}

	return finalRanges
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

func mergeRanges(a, b Range) Range {
	if !isOverlapping(a, b) {
		panic("oops!")
	}

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

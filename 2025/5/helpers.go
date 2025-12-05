package main

import (
	"bufio"
	"os"
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

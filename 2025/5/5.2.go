package main

import (
	"fmt"
)

func hasAnyOverlaps(ranges []Range) bool {
	for _, r1 := range ranges {
		for _, r2 := range ranges {
			if r1.start == r2.start && r1.end == r2.end {
				continue
			}

			if isOverlapping(r1, r2) {
				return true
			}
		}
	}
	return false
}

func processThisRound(ranges []Range) []Range {
	var finalRanges []Range
	for _, r1 := range ranges {
		var overlaps []Range

		for _, r2 := range ranges {
			// skip dupes, we'll add r1 to finalRanges no matter what (either as itself or within the merge)
			if r1.start == r2.start && r1.end == r2.end {
				continue
			}

			if isOverlapping(r1, r2) {
				overlaps = append(overlaps, r2)
			}
		}

		// no overlaps for this one, yay!
		if len(overlaps) == 0 {
			finalRanges = append(finalRanges, r1)
			continue
		}

		// we have overlaps
		overlaps = append(overlaps, r1)
		merged := mergeRanges(overlaps)
		finalRanges = append(finalRanges, merged[0])
	}

	return dedupe(finalRanges)
}

func algorithm5_2(lines []string) int {
	ranges := getRanges(lines)

	fmt.Println("ranges before: ", ranges)
	for hasAnyOverlaps(ranges) {
		ranges = processThisRound(ranges)
		fmt.Println("ranges after: ", ranges)
	}

	total := 0
	for _, r := range ranges {
		total += (r.end - r.start) + 1
	}

	return total
}

package main

import (
	"fmt"
	"slices"
)

func sortRanges(ranges []Range) {
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.start - b.start
	})
}

func getOverlappingRanges(ranges []Range) []Range {
	var overlappingRanges []Range

	for _, r1 := range ranges {
		for _, r2 := range ranges {
			if r1 == r2 {
				continue
			}

			if isOverlapping(r1, r2) {
				overlappingRanges = append(overlappingRanges, r1)
				overlappingRanges = append(overlappingRanges, r2)
			}
		}
	}

	return dedupe(overlappingRanges)
}

func helperAlgorithm5_2_2(ranges []Range) int {
	// get overlapping ranges
	// while there are overlapping ranges, do the following:
	//   1. make a list of newRanges from the overlapping ranges
	//   2. get a set of all the initial ranges (that weren't in overlapping) and the newRanges (from the overlapping)
	//   3. repeat

	var finalRanges []Range
	for _, r := range ranges {
		finalRanges = append(finalRanges, r)
	}
	overlaps := getOverlappingRanges(ranges)
	for len(overlaps) >= 1 {
		fmt.Println("hello jeff! iteration", len(overlaps))
		fmt.Println("overlaps: ", overlaps)
		fmt.Println("finalRanges: ", finalRanges)
		// get overlapping set
		// get nonoverlaps
		// add nonoverlaps to finalRanges
		// then continue with previous logic

		var nonOverlaps []Range
		for _, r := range finalRanges {
			hasOverlap := false
			for _, overlap := range overlaps {
				if r.start == overlap.start && r.end == overlap.end {
					continue
				}
				if isOverlapping(r, overlap) {
					hasOverlap = true
					break
				}
			}
			if !hasOverlap {
				nonOverlaps = append(nonOverlaps, r)
			}
		}

		finalRanges = nil

		var newRanges []Range
		for _, r1 := range overlaps {
			var theseOverlaps []Range
			theseOverlaps = append(theseOverlaps, r1)
			for _, r2 := range overlaps {
				if isOverlapping(r1, r2) {
					theseOverlaps = append(theseOverlaps, r2)
				}
			}
			newRanges = append(newRanges, mergeRangesUntilOne(theseOverlaps)...)
		}
		// finalRangesMap := make(map[int]struct{})
		// for _, r1 := range overlaps {
		// 	_, exists := finalRangesMap[r1.start]
		// 	if exists {
		// 		continue
		// 	}
		// 	finalRangesMap[r1.start] = struct{}{}

		// 	var rangesWithSameStartOrEnd []Range
		// 	for _, r2 := range overlaps {
		// 		if r2.start == r1.start || r2.end == r1.end {
		// 			rangesWithSameStartOrEnd = append(rangesWithSameStartOrEnd, r2)
		// 		}
		// 	}

		// 	newRanges = slices.Concat(newRanges, mergeRangesUntilOne(rangesWithSameStartOrEnd))
		// }

		finalRanges = slices.Concat(nonOverlaps, dedupe(newRanges))
		overlaps = getOverlappingRanges(finalRanges)
	}

	total := 0
	for _, r := range finalRanges {
		total += (r.end - r.start) + 1
	}
	return total
}

func algorithm5_2_2(lines []string) int {
	ranges := getRanges(lines)
	// sortRanges(ranges)

	// var overlappingRanges []Range

	// for _, r1 := range ranges {
	// 	for _, r2 := range ranges {
	// 		if isOverlapping(r1, r2) {
	// 			overlappingRanges = append(overlappingRanges, mergeRanges(r1, r2))
	// 		}
	// 	}
	// }

	// for _, r := range newRanges {
	// 	fmt.Println(r)
	// }

	// fmt.Println("Num Ranges: ", len(ranges))
	// fmt.Println("Num new ranges: ", len(newRanges))

	// total := 0
	// for _, r := range newRanges {
	// 	total += r.end - r.start + 1
	// }
	return helperAlgorithm5_2_2(ranges)
}

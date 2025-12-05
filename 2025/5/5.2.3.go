package main

func isOverlapping3(a, b Range) bool {
	return (a.start <= b.start && b.start <= a.end) || (b.start <= a.start && a.start <= b.end)
}

func algorithm5_2_3(lines []string) int {
	ranges := getRanges(lines)

	freshIds := make(map[int]struct{})

	for _, r := range ranges {
		current := r.start
		for current <= r.end {
			freshIds[current] = struct{}{}
			current += 1
		}
	}

	return len(freshIds)
}

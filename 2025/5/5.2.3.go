package main

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

package problems

func ReportRepair(entries []int) int {
	for i := range entries {
		for j := range entries {
			if entries[i]+entries[j] == 2020 {
				return entries[i] * entries[j]
			}
		}
	}

	return 0
}

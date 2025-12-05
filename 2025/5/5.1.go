package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func algorithm5_1(lines []string) int {
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

	return total
}

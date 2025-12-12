package main

import (
	"fmt"
	"slices"
	"strconv"
)

type Present []string

type Tree struct {
	width         int
	length        int
	presentCounts []int
}

func makePresents(lines []string) []Present {
	var presents []Present

	i := 0
	isPresent := false
	var present Present
	for _, l := range lines {
		if isPresent && l == "" {
			presents = append(presents, slices.Clone(present))
			present = nil
			isPresent = false
			i++
			continue
		}

		if !isPresent && string(l[0]) == strconv.Itoa(i) {
			isPresent = true
			continue
		}

		if isPresent {
			present = append(present, l)
		}
	}

	return presents
}

func algorithm12_1(lines []string) int {
	total := 0

	presents := makePresents(lines)
	fmt.Println("num presents: ", len(presents))
	for _, p := range presents {
		fmt.Println(p)
	}

	return total
}

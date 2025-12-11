package main

import (
	"fmt"
	"slices"
)

func algorithm11_2(lines []string) int {
	total := 0
	deviceMap := makeDeviceMap(lines)
	fmt.Println("deviceMap: ", deviceMap)

	var stack [][]string
	for _, next := range deviceMap["svr"] {
		stack = append(stack, []string{next})
	}

	for len(stack) > 0 {
		var next [][]string
		next = append(next, stack...)
		stack = nil

		for _, current := range next {
			if current[len(current)-1] == "out" {
				dacIndex := slices.Index(current, "dac")
				fftIndex := slices.Index(current, "fft")
				if dacIndex != -1 && fftIndex != -1 {
					total += 1
				}
				continue
			}

			for _, nextDevice := range deviceMap[current[len(current)-1]] {
				stack = append(stack, slices.Concat(current, []string{nextDevice}))
			}
		}
	}

	return total
}

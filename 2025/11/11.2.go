package main

import (
	"fmt"
)

func algorithm11_2(lines []string) int {
	total := 0
	deviceMap := makeDeviceMap(lines)
	fmt.Println("deviceMap: ", deviceMap)

	var stack []string
	stack = append(stack, deviceMap["you"]...)
	for len(stack) > 0 {
		var next []string
		next = append(next, stack...)
		stack = nil

		for _, current := range next {
			if current == "out" {
				total += 1
				continue
			}

			stack = append(stack, deviceMap[current]...)
		}
	}

	return total
}

package main

import (
	"fmt"
	"strings"
)

func algorithm11_2(lines []string) int {
	total := 0
	deviceMap := makeDeviceMap(lines)
	fmt.Println("deviceMap: ", deviceMap)

	var stack []string
	stack = append(stack, deviceMap["svr"]...)
	for len(stack) > 0 {
		var next []string
		next = append(next, stack...)
		stack = nil

		for _, current := range next {
			// fmt.Println("current: ", current)
			currentDevice := current[len(current)-3:]
			if currentDevice == "out" {
				dacIndex := strings.Index(current, "dac")
				fftIndex := strings.Index(current, "fft")
				if dacIndex != -1 && fftIndex != -1 {
					total += 1
				}
				continue
			}

			for _, nextDevice := range deviceMap[currentDevice] {
				stack = append(stack, current+nextDevice)
			}
		}
	}

	return total
}

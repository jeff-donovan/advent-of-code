package main

import (
	"fmt"
	"strings"
)

func addToCache(cache map[string]struct{}, current string) {
	key := strings.Clone(current)
	for strings.Index(key, "dac") != -1 && strings.Index(key, "fft") != -1 {
		cache[key] = struct{}{}
		key = key[:len(key)-3]
	}
}

func algorithm11_2(lines []string) int {
	total := 0
	deviceMap := makeDeviceMap(lines)
	fmt.Println("deviceMap: ", deviceMap)

	// plan
	//  - need some sort of map to keep track of paths that worked AND had dac and fft
	//  - trim the final 3 characters from the path string and add to map as long as both dac and fft are still there
	//  - FIRST - just build the map during the `currentDevice == "out"` block
	//  - SECOND - figure out where to check the map

	cache := make(map[string]struct{})

	var stack []string
	stack = append(stack, deviceMap["svr"]...)
	for len(stack) > 0 {
		var next []string
		next = append(next, stack...)
		stack = nil

		for _, current := range next {
			_, exists := cache[current]
			if exists {
				total += 1
				continue
			}

			currentDevice := current[len(current)-3:]
			if currentDevice == "out" {
				dacIndex := strings.Index(current, "dac")
				fftIndex := strings.Index(current, "fft")
				if dacIndex != -1 && fftIndex != -1 {
					addToCache(cache, currentDevice)
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

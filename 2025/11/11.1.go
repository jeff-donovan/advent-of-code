package main

import (
	"fmt"
	"strings"
)

func makeDeviceMap(lines []string) map[string][]string {
	deviceMap := make(map[string][]string)
	for _, l := range lines {
		key := l[:strings.Index(l, ":")]
		outputString := l[strings.Index(l, ":")+1:]
		outputSplit := strings.Split(outputString, " ")
		var outputs []string
		for _, o := range outputSplit {
			val := strings.Trim(o, " ")
			if val != "" {
				outputs = append(outputs, val)
			}
		}
		deviceMap[key] = outputs
	}
	return deviceMap
}

func algorithm11_1(lines []string) int {
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

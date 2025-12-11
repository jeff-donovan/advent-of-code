package main

import (
	"fmt"
	"strings"
)

type Device struct {
	name    string
	devices []string
}

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
	return total
}

package main

import "fmt"

func makeBackwardsDeviceMap(deviceMap map[string][]string) map[string][]string {
	backwardsMap := make(map[string][]string)

	for start, endNodes := range deviceMap {
		for _, end := range endNodes {
			if value, ok := backwardsMap[end]; ok {
				backwardsMap[end] = addIfNotExists(value, start)
			} else {
				backwardsMap[end] = []string{start}
			}
		}
	}

	return backwardsMap
}

func addIfNotExists(nodes []string, newNode string) []string {
	for _, n := range nodes {
		if n == newNode {
			return nodes
		}
	}

	return append(nodes, newNode)
}

func algorithm11_2_2(lines []string) int {
	total := 0
	deviceMap := makeDeviceMap(lines)
	backwardsMap := makeBackwardsDeviceMap(deviceMap)

	fmt.Println(deviceMap)
	fmt.Println()
	fmt.Println(backwardsMap)

	fmt.Println()
	fmt.Println("deviceMap: ", len(deviceMap))
	fmt.Println("backwardsMap: ", len(backwardsMap))

	return total
}

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

func dfsNodesThatLeadFromOutToSvr(backwardsMap map[string][]string) []string {
	validNodes := make(map[string]struct{})
	currentNode := "out"
	dfs(backwardsMap, currentNode, validNodes)

	var result []string
	for node := range validNodes {
		result = append(result, node)
	}
	return result
}

func dfs(backwardsMap map[string][]string, currentNode string, validNodes map[string]struct{}) bool {
	if currentNode == "svr" {
		return true
	}

	neighbors := backwardsMap[currentNode]
	for _, nextNode := range neighbors {
		isValid := dfs(backwardsMap, nextNode, validNodes)
		if isValid {
			validNodes[nextNode] = struct{}{}
			validNodes[currentNode] = struct{}{}
			return true
		}
	}

	return false
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

	nodesToConsider := dfsNodesThatLeadFromOutToSvr(backwardsMap)
	fmt.Println()
	fmt.Println("nodes to consider: ", len(nodesToConsider))
	fmt.Println(nodesToConsider)

	return total
}

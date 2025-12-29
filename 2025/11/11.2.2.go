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

func dfsNodesThatLeadFromSvrToOut(graph map[string][]string) []string {
	validNodes := make(map[string]struct{})
	currentNode := "svr"
	dfs(graph, currentNode, validNodes)

	var result []string
	for node := range validNodes {
		result = append(result, node)
	}
	return result
}

func dfs(graph map[string][]string, currentNode string, validNodes map[string]struct{}) bool {
	if currentNode == "out" {
		return true
	}

	if _, ok := validNodes[currentNode]; ok {
		return true
	}

	neighbors := graph[currentNode]

	areAnyNeighborsValid := false
	for _, nextNode := range neighbors {
		isValid := dfs(graph, nextNode, validNodes)
		if isValid {
			areAnyNeighborsValid = true
		}
	}

	if areAnyNeighborsValid {
		validNodes[currentNode] = struct{}{}
	}

	return areAnyNeighborsValid
}

func algorithm11_2_2(lines []string) int {
	total := 0
	deviceMap := makeDeviceMap(lines)

	// fmt.Println(deviceMap)
	// fmt.Println()

	fmt.Println("deviceMap: ", len(deviceMap))

	nodesToConsider := dfsNodesThatLeadFromSvrToOut(deviceMap)
	fmt.Println()
	fmt.Println("nodes to consider: ", len(nodesToConsider))
	// fmt.Println(nodesToConsider)

	return total
}

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

func dfsCountPaths(graph map[string][]string) int {
	paths := make(map[string]int)
	currentNode := "svr"
	endNode := "out"
	return dfs(graph, paths, currentNode, endNode)
}

func dfs(graph map[string][]string, paths map[string]int, currentNode string, endNode string) int {
	if currentNode == endNode {
		return 1
	}

	if numPaths, ok := paths[currentNode]; ok {
		return numPaths
	}

	neighbors := graph[currentNode]

	numPathsFromNeighbors := 0
	for _, nextNode := range neighbors {
		numPathsFromNeighbors += dfs(graph, paths, nextNode, endNode)
	}

	paths[currentNode] = numPathsFromNeighbors

	return numPathsFromNeighbors
}

func algorithm11_2_2(lines []string) int {
	// total := 0
	deviceMap := makeDeviceMap(lines)

	// fmt.Println(deviceMap)
	// fmt.Println()

	fmt.Println("deviceMap: ", len(deviceMap))

	return dfsCountPaths(deviceMap)
	// fmt.Println()
	// fmt.Println("nodes to consider: ", len(nodesToConsider))
	// // fmt.Println(nodesToConsider)

	// return total
}

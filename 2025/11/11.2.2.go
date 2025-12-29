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

func dfsCountPathsThatLeadFromSvrToOut(graph map[string][]string) int {
	scores := make(map[string]int)
	path := "svr"
	return dfs(graph, scores, path)
}

func dfs(graph map[string][]string, scores map[string]int, path string) int {
	currentNode := path[len(path)-3:]
	if currentNode == "out" {
		return 1
	}

	if numPaths, ok := scores[currentNode]; ok {
		return numPaths
	}

	neighbors := graph[currentNode]

	numPathsFromNeighbors := 0
	for _, nextNode := range neighbors {
		nextPath := path + nextNode
		numPathsFromNeighbors += dfs(graph, scores, nextPath)
	}

	scores[currentNode] = numPathsFromNeighbors

	return numPathsFromNeighbors
}

func algorithm11_2_2(lines []string) int {
	// total := 0
	deviceMap := makeDeviceMap(lines)

	// fmt.Println(deviceMap)
	// fmt.Println()

	fmt.Println("deviceMap: ", len(deviceMap))

	return dfsCountPathsThatLeadFromSvrToOut(deviceMap)
	// fmt.Println()
	// fmt.Println("nodes to consider: ", len(nodesToConsider))
	// // fmt.Println(nodesToConsider)

	// return total
}

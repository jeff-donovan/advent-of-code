package main

func dfsCountPaths(graph map[string][]string) int {
	numPathsFromSvrToFft := dfs(graph, make(map[string]int), "svr", "fft")
	numPathsFromFftToDac := dfs(graph, make(map[string]int), "fft", "dac")
	numPathsFromDacToOut := dfs(graph, make(map[string]int), "dac", "out")
	return numPathsFromSvrToFft * numPathsFromFftToDac * numPathsFromDacToOut
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
	deviceMap := makeDeviceMap(lines)

	return dfsCountPaths(deviceMap)
}

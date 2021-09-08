package main

import (
	"fmt"
)

func convertEdgesToAdjacencyList(edges [][]string) map[string][]string {

	convertedAdjList := make(map[string][]string)

	for _, edgePair := range edges {
		convertedAdjList[edgePair[0]] = append(convertedAdjList[edgePair[0]], edgePair[1])
		convertedAdjList[edgePair[1]] = append(convertedAdjList[edgePair[1]], edgePair[0]) // as this is undirected
	}
	return convertedAdjList
}

func hasPathCyclic(graph map[string][]string, source string, destination string, visited map[string]bool) bool {

	if source == destination {
		return true
	}
	if visited[source] == true {
		return true
	}
	visited[source] = true

	for _, neighbour := range graph[source] {

		if hasPathCyclic(graph, neighbour, destination, visited) {
			return true
		}
	}

	return false
}

func countConnectComponents(graph map[string][]string) int {

	visited := make(map[string]bool)
	count := 0

	for key, _ := range graph {
		if explore(graph, key, visited) == true {
			count++
		}
	}

	return count
}

func explore(graph map[string][]string, node string, visited map[string]bool) bool {

	if visited[node] == true {
		return false
	}
	visited[node] = true
	for _, neighbour := range graph[node] {
		explore(graph, neighbour, visited)
	}
	return true

}

func largestComponent(graph map[string][]string) int {
	visited := make(map[string]bool)
	largestComponent := 0

	for key, _ := range graph {
		size := explorewithCount(graph, key, visited)

		if size > largestComponent {
			largestComponent = size
		}
	}

	return largestComponent
}

func explorewithCount(graph map[string][]string, node string, visited map[string]bool) int {

	
	if visited[node] == true {
		return 0
	}
	visited[node] = true
	size := 1
	for _, neighbour := range graph[node] {
		size += explorewithCount(graph, neighbour, visited)
	}
	return size
}

func main() {
	edges := [][]string{{"i", "j"}, {"k", "i"}, {"m", "k"}, {"k", "l"}, {"o", "n"}}
	fmt.Println(edges)
	fmt.Println("----")
	adjacencyList := convertEdgesToAdjacencyList(edges)
	fmt.Println(adjacencyList)
	fmt.Println("----")

	visited := make(map[string]bool)

	fmt.Println(hasPathCyclic(adjacencyList, "i", "k", visited))

	fmt.Println("----")

	fmt.Println(countConnectComponents(adjacencyList))

	fmt.Println("----")
	fmt.Println(largestComponent(adjacencyList))
}

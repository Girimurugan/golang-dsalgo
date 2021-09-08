package main

import (
	"fmt"
)

func hasPath(graph map[string][]string, source string, destination string)bool{

	// base condition
	if source == destination{
		return true
	}

	// recursion

	for _, neighbour := range graph[source]{

		if hasPath(graph,neighbour,destination){
			return true
		}
	}

	return false
}


func main(){

	graphAdjacencyList := make(map[string][]string)

	graphAdjacencyList["a"] = []string{"b", "c"}
	graphAdjacencyList["b"] = []string{"d"}
	graphAdjacencyList["c"] = []string{"e"}
	graphAdjacencyList["d"] = []string{"f"}
	graphAdjacencyList["e"] = []string{}
	graphAdjacencyList["f"] = []string{}
	fmt.Println(graphAdjacencyList)

	fmt.Println(hasPath(graphAdjacencyList,"a","f"))

}
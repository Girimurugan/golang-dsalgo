package main

import (
	"fmt"
)


type stack struct {
	nodes []string
}

// implement a push function

func (s *stack) push(node string) {
	s.nodes = append(s.nodes, node)
}

// implement a pop function
func (s *stack) pop() string {
	toPop := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]

	return toPop
}


func depthFirstSearchRecursive(graph map[string][]string, source string){
	fmt.Println(source)

	for _,neighbour := range  graph[source]{
		depthFirstSearchRecursive(graph,neighbour)
	}
}

func depthFirstSearchIterative(graph map[string][]string, source string){

stack := &stack{}

stack.push(source)

for len(stack.nodes) > 0{

	currentNode := stack.pop()
	fmt.Println(currentNode)

	for _, neighbour := range graph[currentNode]{
		stack.push(neighbour)
	}

}

}


func main(){

	graphAdjacencyList := make(map[string][]string)

	graphAdjacencyList["a"] = []string{"b","c"}
	graphAdjacencyList["b"] = []string{"d"}
	graphAdjacencyList["c"] = []string{"e"}
	graphAdjacencyList["d"] = []string{"f"}
	graphAdjacencyList["e"] = []string{}
	graphAdjacencyList["f"] = []string{}
	fmt.Println(graphAdjacencyList)

	depthFirstSearchIterative(graphAdjacencyList,"a")

	fmt.Println("")

	depthFirstSearchRecursive(graphAdjacencyList,"a")

}
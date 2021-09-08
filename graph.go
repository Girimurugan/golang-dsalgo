package main

import "fmt"

// Definition for the Graph
type Graph struct {
	nodes []*Nodes
}

// Definition for the Nodes/Vertexs
type Nodes struct {
	key       int
	adjacents []*Nodes
}

// Add an Node
func (g *Graph) AddNode(key int) {

	if g.contains(key) != true {
		g.nodes = append(g.nodes, &Nodes{key: key})
	}
}

// check node is already existing

func (g *Graph) contains(key int) bool {

	for _, v := range g.nodes {
		if v.key == key {
			return true
		}
	}
	return false
}

//print the Graph as Adjacency List

func (g *Graph) Print() {

	for _, v := range g.nodes {

		fmt.Printf("\n Node: %v ---> ", v.key)

		for _, a := range v.adjacents {
			fmt.Printf("%v ", a.key)
		}
	}

}

// Add an Edge

func (g *Graph) AddEdge(fromKey int, toKey int){

	if fromKey != 0 && toKey != 0{

		fromNode := g.getNode(fromKey)
		toNode := g.getNode(toKey)

		fromNode.adjacents = append(fromNode.adjacents, toNode)


	}

}

// Get the node for a given key
 
func (g *Graph) getNode(key int) *Nodes{

	if key != 0 {

		for _ ,v := range g.nodes{
			if v.key == key {
				return v
			}
		}

	}
	return &Nodes{}
}

// check if the Edge does not exist already

func main() {

	myGraph := &Graph{}

	myGraph.AddNode(10)
	myGraph.AddNode(20)
	myGraph.AddNode(30)
	myGraph.AddNode(40)

	myGraph.AddEdge(10,20)
	myGraph.AddEdge(10,30)
	myGraph.AddEdge(10,40)
	myGraph.AddEdge(20,40)
	myGraph.AddEdge(30,20)

	myGraph.Print()

}

package main

import (
	"fmt"
)

type queue struct {
	items []string
}

func (q *queue) enqueue(node string) {
	q.items = append(q.items, node)
}

func (q *queue) dequeue() string {

	toDequeue := q.items[0]
	q.items = q.items[1:]

	return toDequeue
}

func breadthFirstPrintInOrder(graph map[string][]string, source string) {

	queue := &queue{}

	currentNode := source
	queue.enqueue(currentNode)

	for len(queue.items) > 0 {

		currentNode = queue.dequeue()
		fmt.Println(currentNode)

		for _, neighbour := range graph[currentNode] {

			queue.enqueue(neighbour)

		}
	}

}

func breadthFirstPrintPreOrder(graph map[string][]string, source string) {

	queue := &queue{}

	currentNode := source
	queue.enqueue(currentNode)

	for len(queue.items) > 0 {

		currentNode = queue.dequeue()
		

		for _, neighbour := range graph[currentNode] {

			queue.enqueue(neighbour)

		}
		fmt.Println(currentNode)
	}

}

func main() {

	graphAdjacencyList := make(map[string][]string)

	graphAdjacencyList["a"] = []string{"b", "c"}
	graphAdjacencyList["b"] = []string{"d"}
	graphAdjacencyList["c"] = []string{"e"}
	graphAdjacencyList["d"] = []string{"f"}
	graphAdjacencyList["e"] = []string{}
	graphAdjacencyList["f"] = []string{}
	fmt.Println(graphAdjacencyList)

	breadthFirstPrintInOrder(graphAdjacencyList,"a")
	fmt.Println("----")
	breadthFirstPrintPreOrder(graphAdjacencyList,"a")

}

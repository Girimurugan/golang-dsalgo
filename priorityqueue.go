package main

import (
	"fmt"
)

// item definition for the queue item
type Item struct{
	value string
	priority int
}

// definition of Priority Queue
type PriorityQueue struct{
	items []Item
}

// Functions to work on the Priority queue as a HEAP

// utility functions

func parent(index int) int{
return (index-1)/2
}

func left(index int) int{
	return (index * 2) + 1
}

func right(index int) int{
	return (index * 2) + 2
}

func (pq *PriorityQueue) swap(index1 int, index2 int){
	pq.items[index1], pq.items[index2] = pq.items[index2], pq.items[index1]
}


// insert into the Priority queue (2 steps - append & heapify up)

func (pq *PriorityQueue) Insert(item Item){

	pq.items = append(pq.items, item)

	pq.maxHeapifyUp(len(pq.items)-1)

}

func (pq *PriorityQueue) maxHeapifyUp(startIndex int){

	for pq.items[startIndex].priority > pq.items[parent(startIndex)].priority{

		pq.swap(startIndex, parent(startIndex))

		startIndex = parent(startIndex)


	}

}

// Extract Function (Step 1 - Return Index 0 , Step 2 - Promate the last leaf to root and Step 3 - Heapify Down)
func (pq *PriorityQueue) Extract() Item{

extractedItem := pq.items[0]

pq.items[0] = pq.items[len(pq.items)-1]
pq.items = pq.items[:len(pq.items)-1]

pq.maxHeapifyDown()

return extractedItem

}


func (pq *PriorityQueue) maxHeapifyDown(){

	index :=  0

	comparewithChild := 0

	lastIndex := len(pq.items)-1


	l := left(index)
	r := right(index)

	for l <= lastIndex{

		if (l == lastIndex){
			comparewithChild = l
		}else if (pq.items[l].priority > pq.items[r].priority){
			comparewithChild = l

		}else{
			comparewithChild = r
		}

		if(pq.items[index].priority < pq.items[comparewithChild].priority){
			pq.swap(index,comparewithChild)
			index = comparewithChild
			l = left(index)
			r = right(index)

		}else{
			return	
		}

	}
}


func main(){

priorityQueue := &PriorityQueue{}



newItem := Item{value:"Aradhana",priority:8} 
priorityQueue.Insert(newItem)

newItem = Item{value:"Preethi",priority:32} 
priorityQueue.Insert(newItem)

newItem = Item{value:"Giri",priority:40} 
priorityQueue.Insert(newItem)

newItem = Item{value:"Sri Sai",priority:10} 
priorityQueue.Insert(newItem)


fmt.Println(priorityQueue.Extract())

fmt.Println(priorityQueue.Extract())

fmt.Println(priorityQueue.Extract())

fmt.Println(priorityQueue.Extract())

}
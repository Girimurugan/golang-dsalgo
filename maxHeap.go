package main

import "fmt"

// define the structure for the heap
type Heap struct {
	values []int
}

// Insert into the Heap
func (h *Heap) Insert(key int) {

	// Add the new key at the end
	h.values = append(h.values, key)

	// Max Heapify Up
	h.maxHeapifyUp(len(h.values) - 1)
}

func (h *Heap) maxHeapifyUp(index int) {

	for h.values[index] > h.values[parent(index)] {

		//swap the values
		h.swap(parent(index), index)
		index = parent(index)
	}

	//compare the value with the parent , if the value in index is greater than parent then swap it

}

func (h *Heap) swap(index1 int, index2 int) {
	h.values[index1], h.values[index2] = h.values[index2], h.values[index1]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (index * 2) + 1
}

func right(index int) int {
	return (index * 2) + 2
}

// Extract into the heap
func (h *Heap) Extract() int {

	// get the top of the heap
	extractedValue := h.values[0]

	length := len(h.values) - 1
	// move the last element first
	h.values[0] = h.values[length]

	// shorten the slice
	h.values = h.values[:length]

	// heapifydown
	h.maxHeapifyDown(0)

	// return the index 0

	return extractedValue
}

func (h *Heap) maxHeapifyDown(index int) {

	lastIndex := len(h.values) - 1

	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {

		if l == lastIndex { //last node
			childToCompare = l
		} else if h.values[l] > h.values[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.values[index] < h.values[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

func main() {

	mh := &Heap{}
	buildHeap := []int{10, 30, 40, 7, 8, 56, 34}

	for _, v := range buildHeap {
		mh.Insert(v)
		fmt.Println(mh)
	}

	fmt.Println(mh.Extract())
	fmt.Println(mh)

	fmt.Println(mh.Extract())
	fmt.Println(mh)

	fmt.Println(mh.Extract())
	fmt.Println(mh)
}

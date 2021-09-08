package main

import (
	"fmt"
)

func quickSort(array []int, start int, end int) []int {

	if start < end {

		p := 0

		array, p := parition(array, start, end)
		array = quickSort(array, start, p-1)
		array = quickSort(array, p+1, end)

	}
	return array

}

func parition(array []int, start int, end int) ([]int, int) {

	pivot := array[end]

	i := start

	for j := start; j < end; j++ {

		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}

	array[i], array[end] = array[end], array[i]
	return array, i
}

func main() {

	testArray := []int{4, 7, 2, 4, 8, 2, 4, 2}

	fmt.Println(testArray)
	fmt.Println(quickSort(testArray, 0, len(testArray)-1))

}

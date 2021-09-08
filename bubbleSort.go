package main

import "fmt"

// function for bubble sort

func bubbleSort(input []int) {

	length := len(input)
	for i := 0; i < length-1; i++ {

		for j := 0; j < length-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}

}

func main() {

	testArray := []int{3, 6, 1, 2, 9, 5}
	fmt.Println(testArray)

	bubbleSort(testArray)
	fmt.Println(testArray)
}

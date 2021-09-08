package main

import "fmt"

// Function for binary search

func binarySearch(input []int, searchable int, start int, end int) int {

// Base condition
	   if start > end{
		   return -1
	   }
		// find the median
		median := (start+end)/2

		// check if the median is searcable element, then return the index
		if input[median] == searchable {
			return median
		} else if input[median] > searchable {
			return binarySearch(input, searchable, start, median-1)
		} else if input[median] < searchable{
			return binarySearch(input, searchable, median+1, end)
		}
	
	//If not recurse for on the right for if the element in the is smallee
	return -1
}

func main() {

	testArray := []int{10, 45, 58, 68, 76, 89}

	length := len(testArray)
	fmt.Println(binarySearch(testArray, 45, 0, length-1))

	fmt.Println(binarySearch(testArray, 89, 0, length-1))

	fmt.Println(binarySearch(testArray, 67, 0, length-1))

}

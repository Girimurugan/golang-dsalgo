package main

import "fmt"

func mergeSort(input []int) []int {

	// find the lenght
	length := len(input)

	// Entry condition checks
	if length < 2 {
		return input
	}

	// split into 2 and later merge
	first := mergeSort(input[:length/2])
	second := mergeSort(input[length/2:])
	return merge(first, second)

}

func merge(first []int, second []int) []int {

	final := []int{}
	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}
	for ; i < len(first); i++ {
		final = append(final, first[i])
	}
	for ; j < len(second); j++ {
		final = append(final, second[j])
	}

	return final
}

func main() {

	testArray := []int{4, 7, 2, 4, 8, 2, 4, 2}

	fmt.Println(testArray)
	fmt.Println(mergeSort(testArray))
	fmt.Println(testArray)

}

package main

import (
	"fmt"
)

func insertionSort(input []int){

	for i:=0; i< len(input); i++{

		key := input[i]

		j := i-1

		for (j >= 0 && input[j]>key){
			input[j+1] = input[j]
			j = j-1
		}
		input[j+1] = key

	}

}


func main(){

	testArray := []int{3, 6, 1, 2, 9, 5}
	fmt.Println(testArray)

	insertionSort(testArray)
	fmt.Println(testArray)

}
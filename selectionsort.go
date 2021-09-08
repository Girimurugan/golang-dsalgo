package main

import (
	"fmt"
)


func selectionSort(input []int){

length := len(input)

for i:=0 ; i< length-1 ; i++{

	for j:=i; j < length; j++{

		if input[i] > input[j]{
			input[i],input[j] = input[j], input[i]
		}

	}


}


}

func main(){

	testArray := []int{3, 6, 1, 2, 9, 5}
	fmt.Println(testArray)

	selectionSort(testArray)
	fmt.Println(testArray)


}
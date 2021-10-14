package main

import "fmt"

func quickSort2(array *[]int, start int, end int){

if start < end {

	pivot := parition(array, start, end)

	quickSort2(array, start, pivot-1)
	quickSort2(array, pivot+1, end)
}


}

func parition(array *[]int, start int, end int)int{


	pivot := (*array)[end]

	i := start

	for j:= start ; j < end ; j++{

		if (*array)[j] < pivot{

			(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
			i++
		}

	}

	(*array)[i], (*array)[end] =  (*array)[end], (*array)[i]
	return i 


}

func main(){
	testArray := []int{4, 7, 2, 4, 8, 2, 4, 2}

	fmt.Println(testArray)
	quickSort2(&testArray, 0, len(testArray)-1)

	fmt.Println(testArray)

}
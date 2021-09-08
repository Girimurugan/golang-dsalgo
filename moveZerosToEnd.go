package main

import "fmt"

// Create a Function to move Zeros to end in an array

func moveZerosToEndBrute(input []int) []int{
	length := len(input)
	output := make([]int, length)

	j := 0
	i := 0
	for i = 0; i < len(input); i++ {

		if input[i] != 0 {
			output[j] = input[i]
			j++
		}

	}

	for j < i{
		output[j] = 9
		j++
	}
	return output
}

func moveZeroToEnd(input []int) []int{

	count := 0
	for i:= 0 ; i< len(input); i++{
		if input[i] != 0{
			input[count] = input[i]
			count++
		}
	}

	for count != len(input){
		input[count] = 0
		count++
	}
	return input
}

// call the function and test it

func main(){

	input := [...]int{2,0,3,4,0,6,7,0}
	input2 := [...]int{2,3}
	fmt.Println(moveZerosToEndBrute(input[:]))
	fmt.Println(moveZerosToEndBrute(input2[:]))

	fmt.Println(moveZeroToEnd((input[:])))
	fmt.Println(moveZeroToEnd((input2[:])))
}
package main

import "fmt"

// Structure for the stack

type stack struct{
	items []int
}

// Push Function

func (s *stack)push(value int){
	s.items = append(s.items, value)
}

// Pop Function
func (s *stack)pop() int{

	length := len(s.items)-1
	if length >= 0{
	popedValued := s.items[length]
	s.items = s.items[:length]
	

	return	popedValued
	}else{
		return -1
	}
}




func main(){

	mystack := &stack{}

	mystack.push(12)
	mystack.push(13)
	mystack.push(14)

	fmt.Println(mystack.pop())
	fmt.Println(mystack.pop())
	fmt.Println(mystack.pop())
	fmt.Println(mystack.pop())

}
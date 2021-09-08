package main

import "fmt"

// define the struct for stack

type stack struct {
	items []int
}

// implement a push function

func (s *stack) push(item int) {
	s.items = append(s.items, item)
}

// implement a pop function
func (s *stack) pop() int {
	toPop := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return toPop
}

func main() {

	myStack := &stack{}

	myStack.push(10)
	myStack.push(30)
	myStack.push(40)
	myStack.push(60)

	fmt.Println(myStack)

	fmt.Println(myStack.pop())
	fmt.Println(myStack)

}

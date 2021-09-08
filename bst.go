package main

import "fmt"

// Define the type for a binary tree node
type node struct {
	value int
	left  *node
	right *node
}

// functional to add a node into BST
func (n *node) insert(newValue int) {

	// if the new value is lesser than parent then move left
	if n.value > newValue {

		if n.left == nil {
			newLeftNode := &node{newValue, nil, nil}
			n.left = newLeftNode

		} else {
			n.left.insert(newValue)
		}

	} else if n.value < newValue {
		// if the new value is more than parent move right

		if n.right == nil {
			newRightNode := &node{newValue, nil, nil}
			n.right = newRightNode

		} else {
			n.right.insert(newValue)
		}

	}
}



// Search for value in BST
func (n *node) search(searchFor int) bool {

	if n == nil {
		return false
	}
	if n.value > searchFor {

		return n.left.search(searchFor)

	} else if n.value < searchFor {

		return n.right.search(searchFor)

	}

	return true
}

func (n *node) printInOrder(){

	if n == nil {
		return
	}
	n.left.printInOrder()
	fmt.Println(n.value)
	n.right.printInOrder()
}

func (n *node) printPreOrder(){

	if n == nil {
		return
	}
	fmt.Println(n.value)
	n.left.printPreOrder()
	n.right.printPreOrder()
}

func (n *node) printPostOrder(){

	if n == nil {
		return
	}
	
	n.left.printPostOrder()
	n.right.printPostOrder()
	fmt.Println(n.value)
}

func main() {

	rootNode := &node{100, nil, nil}
	rootNode.insert(50)
	rootNode.insert(20)
	rootNode.insert(30)
	rootNode.insert(40)
	fmt.Println(rootNode)

	fmt.Println(rootNode.search(50))
	fmt.Println(rootNode.search(350))


	fmt.Println("-----")
	rootNode.printInOrder()

	fmt.Println("-----")
	rootNode.printPreOrder()

	fmt.Println("-----")
	rootNode.printPostOrder()
}

package main

import (
	"fmt"
)

// Type defintion for node

type node struct {
	value int
	next  *node
}

// type defintiion for a linked list
type linkedList struct {
	head   *node
	length int
}

// function to pre-append a node

func (ll *linkedList) preapend(data int) {

	existingHead := ll.head
	newNode := node{value: data}
	ll.head = &newNode
	ll.head.next = existingHead
	ll.length++
}

// function to print a linked list
func (ll *linkedList) printLinkedList() {

	printNode := ll.head
	fmt.Print(printNode.value)

	for printNode.next != nil {
		printNode = printNode.next
		fmt.Print(" -> ")
		fmt.Print(printNode.value)
		
	
	}

}

// function to delete the node

func (ll *linkedList) deleteNode(data int) {

	currentNode := ll.head

	// if head node is the node to be deleted
	if currentNode.value == data {
		ll.head = currentNode.next
	} else {

		previousNode := currentNode
		currentNode = currentNode.next
		for currentNode.next != nil {

			if currentNode.value == data {
				previousNode.next = currentNode.next
			}
			currentNode = currentNode.next

		}
	}

}

// function to revese the linked list
func (ll *linkedList) reverseLinkedList(){

previousNode := &node{}
currentNode := &node{}
nextNode := &node{}

currentNode = ll.head

for (currentNode != nil){

	nextNode = currentNode.next
	currentNode.next = previousNode
	previousNode = currentNode
	currentNode = nextNode
	

}

ll.head = previousNode

}

func main() {

	myLinkedList := &linkedList{}
	myLinkedList.preapend(32)
	myLinkedList.preapend(20)
	myLinkedList.preapend(45)

	fmt.Println(myLinkedList)
	myLinkedList.printLinkedList()
	myLinkedList.deleteNode(20)
	fmt.Println()

	myLinkedList.preapend(67)
	myLinkedList.preapend(12)
	myLinkedList.printLinkedList()

	fmt.Println()

	fmt.Println("After reversal")
	myLinkedList.reverseLinkedList()
	myLinkedList.printLinkedList()

}

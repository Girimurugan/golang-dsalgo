package main

import "fmt"

//define the struct for queue
type queue struct{
	items []int
}

//enqueue

func (q *queue)enqueue(item int){
	q.items = append(q.items, item)
}

//dequeue
func (q *queue)dequeue()int{
	toDequeue := q.items[0]
	q.items = q.items[1:]

	return toDequeue
}

func main(){

	myQueue := &queue{}

	myQueue.enqueue(23)
	myQueue.enqueue(33)
	myQueue.enqueue(43)

	fmt.Println(myQueue)
	fmt.Println(myQueue.dequeue())
	fmt.Println(myQueue)
}
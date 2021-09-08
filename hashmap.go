package main

import "fmt"



const ArraySize = 5

// define the struct for the hash map
type hashMap struct{
	items [ArraySize]*bucket
}


// define the struct for the bucket
type bucket struct{
	head *bucketNode
}

// define the struct for the bucket node
type bucketNode	struct{
	key string
	nextNode *bucketNode
}

// initialize function
func createHashTable()*hashMap{

	returnHashMap := &hashMap{}

	for i := range returnHashMap.items{
		returnHashMap.items[i] = &bucket{}
	}

	return returnHashMap	

}

// function to Insert into the hash 
func (h *hashMap)insert(key string){

	index := hash(key)
	h.items[index].insertIntoBucket(key)

}

// function to Search
func (h *hashMap) search(key string) bool{
	index := hash(key)
	h.items[index].searchInBucket(key)
	return true
}

// function to Delete
func (h *hashMap)delete (key string) {

}

// Insert into the Bucket
func (b *bucket)insertIntoBucket(key string){
	newNode := &bucketNode{key:key}
	newNode.nextNode = b.head
	b.head = newNode
}

// Search in the Bucket
func (b *bucket) searchInBucket(key string)bool{
	currentNode := b.head

	for currentNode != nil{
		if currentNode.key == key{
			return true
		}
		currentNode = currentNode.nextNode
	}
	return false
}

// Delete in the Bucket

// hashmap function
func hash(key string)int{

	sum := 0
	for _,v := range key{
		sum+=int(v)
	}
	return sum%ArraySize
}

func main(){
 fmt.Println ("Compiled")
}
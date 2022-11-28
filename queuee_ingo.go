package main

import (
	"fmt"
)

type Queue []int

// IsEmpty: check if Queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Push a new value onto the Queue
func (q *Queue) Enqueue(str int) {
	*q = append(*q, str) // Simply append the new value to the end of the queue
}

// Remove and return top element of queue. Return false if Queue is empty.
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	} else {
		element := (*q)[0] // Index into the slice and obtain the element.
		*q = (*q)[1:]      // Remove it from the queue by slicing it off.
		
		return element, true
		

	}
}
func (q *Queue) top() (int, bool){
		if q.IsEmpty() {
			return 0, false
	}else{
		d:= *q
		fmt.Println(d[len(d)-1])
		return 0, true
 	}
}
func main() {
	var queue Queue // create a queue variable of type queue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(35)
	fmt.Println("----------")
	queue.top()
	fmt.Println("---peak or top-------")
	queue.Dequeue()
	queue.Enqueue(45)
	queue.Dequeue()
	queue.top()
	fmt.Println("----------")
	for len(queue) > 0 {
		x, y := queue.Dequeue()
		if y == true {
			fmt.Println(x)
		}
	}
	
}

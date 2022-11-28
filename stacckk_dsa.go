package main

import (
	"fmt"
)

type Stack []int

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str int) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		
		return element, true
		

	}
}
func (s *Stack) Peak() (int, bool){
		if s.IsEmpty() {
			return 0, false
	}else{
		d:= *s
		fmt.Println(d[len(d)-1])
		return 0, true
 	}
}
func main() {
	var stack Stack // create a stack variable of type Stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(35)
	fmt.Println("----------")
	stack.Peak()
	fmt.Println("---peak or top or poped-------")
	stack.Pop()
	stack.Push(45)
	stack.Peak()
	fmt.Println("----------")
	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
	
}

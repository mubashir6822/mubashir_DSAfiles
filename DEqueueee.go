package main

import "fmt"

type node struct {
	data int
	next *node
}
type DEqueue struct {
	len  int
	head,tail *node
	
}

func (q *DEqueue) IsEmpty() bool {
	if q.len == 0 {
		return true
	} else {
		return false
	}
}
func (q *DEqueue) addfirst(e int) {
	new := &node{e, nil}
	if q.head == nil {
		q.head = new
		q.tail = new
		q.len++
	} else {
		new.next = q.head
		q.head = new
		q.len++
	}
}
func (q *DEqueue) addlast(e int) {
	new := &node{e, nil}
	if q.head == nil {
		q.head = new
		q.tail = new
	} else {
		q.tail.next = new
		q.len++
	}
}

func (q *DEqueue) removefirst() {
	if q.head == nil {
		fmt.Println("DEqueuelist is empty")
		return
	}
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
		q.len--
	}

}
func (q *DEqueue) removelast() {
	if q.head == nil {
		fmt.Println("DEqueuelist is empty")
		return
	}
	q.tail.next = nil
	q.len--

}
func (q *DEqueue) first() {
	if q.IsEmpty(){
		fmt.Println("empty")

	}else{
		fmt.Println(q.head)
	}
	
}
func (q *DEqueue) last() {
	fmt.Println(q.tail)
}

func (q *DEqueue) display() {
	p := q.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next

	}
}
func main() {
	dque := DEqueue{}
	dque.addlast(30)
	dque.display()
	fmt.Println()
	dque.addfirst(20)
	dque.display()
	fmt.Println()
	dque.addfirst(10)
	dque.display()
	fmt.Println()

	dque.removefirst()
	dque.display()
	fmt.Println()
	dque.addlast(40)
	dque.display()
	fmt.Println()
	dque.removelast()
	dque.display()
	fmt.Println("\n--------------------------------")
	dque.first()
	fmt.Println("----------------------------------------------------------")
	dque.last()
}

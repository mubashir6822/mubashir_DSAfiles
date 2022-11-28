package main

import "fmt"

type node struct {
	data int
	next *node
}
type singly struct {
	len int
	head *node
	tail *node
}

func (s *singly) IsEmpty() bool {
	if s.len == 0 {
		return true
	} else {
		return false
	}
}

func (s *singly) addfirst(e int) {
	new := &node{e, nil}
	if s.head == nil {
		s.head = new
		s.tail = new
		s.len++
	} else {
		new.next = s.head
		s.head = new
		s.len++
	}
}
func (s *singly) addlast(e int) {
	new := &node{e, nil}
	if s.head == nil {
		s.head = new
		s.tail = new
	} else {
		s.tail.next = new
		s.len++
	}
}
func (s *singly) addany(e, pos int) {
	new := &node{e, nil}
	p := s.head
	i := 0
	for i < pos-1 {
		p = p.next
		i++
	}
	new.next = p.next
	p.next = new
	s.len++
}

func (s *singly) removefirst() {
	if s.head == nil {
		fmt.Println("singlylist is empty")
		return
	}
	s.head = s.head.next
	if s.head == nil {
		s.tail = nil
		s.len--
	}

}
func (s *singly) removelast() {
	if s.head == nil {
		fmt.Println("singlylist is empty")
		return
	}
	s.tail.next = nil
	s.len--

}
func (s *singly) removeany(pos int) {
	p := s.head
	i := 0
	for i < pos-1 {
		p = p.next
		i++
	}
	p.next = p.next.next
	s.len--

}
func (s *singly) display() {
	p := s.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next

	}
}
func main() {
	singll := singly{}
	singll.addlast(30)
	singll.display()
	fmt.Println()
	singll.addfirst(20)
	singll.display()
	fmt.Println()
	singll.addfirst(10)
	singll.display()
	fmt.Println()
	singll.addany(25, 2)
	singll.display()
	fmt.Println()
	singll.removefirst()
	singll.display()
	fmt.Println()
	singll.addlast(40)
	singll.display()
	fmt.Println()
	singll.removelast()
	singll.display()
	fmt.Println()
	singll.removeany(1)
	singll.display()
	fmt.Println()
}

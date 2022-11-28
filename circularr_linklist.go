package main
import "fmt"

type Node struct{
  element int
  next *Node
  
}

type cll struct{
  head *Node
  tail *Node
  size int
}

func (l *cll) len() int{
  return l.size
}

func (l *cll) isempty() bool {
  return l.size==0
}

func (l *cll) addlast(e int){
  new:=&Node{element:e}
  if l.isempty(){
    l.head=new
    l.head.next=new
  }else{
    l.tail.next=new
    new.next=l.head
  }
  l.tail=new
  l.size++
}

func (l *cll) addfirst(e int){
  new := &Node{element:e}
  if l.isempty(){
    l.head=new
    l.head.next=new
    l.tail=new
  }else{
    new.next=l.head
    l.tail.next=new
    l.head=new
  }
  l.size++
}

func (l *cll) addany(e int, position int){
  new:= &Node{element:e}
  if l.isempty(){
    l.head=new
    l.tail=new
  }
  i:=1
  p:=l.head
  for i<position-1{
    p=p.next
    i++
  }
  new.next=p.next
  p.next=new
  l.size++
}

func (l *cll) display(){
  p:=l.head
  if l.isempty(){return}
  i:=0
  for i<l.len(){
    fmt.Print(p.element, "-->")
    p=p.next
    i++
  }
}

func (l *cll) removelast() int{
  if l.isempty(){
    fmt.Println("\n CLL is empty")
    return 0
  }
  e:=l.tail.element
  i:=1
  p:=l.head
  for i<l.len()-1{
    p=p.next
    i++
  }
  p.next=l.head
  l.tail=p
  l.size--
  return e
}

func (l *cll) removefirst() int{
  if l.isempty(){
    fmt.Println("\n CLL is empty")
    return 0
  }
  e:=l.head.element
  l.tail.next=l.head.next
  l.head=l.head.next
  l.size--
  return e
}

func (l *cll) removeany(position int) int{
  if l.isempty(){
    fmt.Println("\n CLL is empty")
    return 0
  }
  i:=1 
  p:=l.head
  for i<position-1{
    p=p.next
    i++
  }
  e:=p.next.element
  p.next=p.next.next
  l.size--
  return e
}

func (l *cll) search(e int) int{
  p:=l.head
  i:=0
  for i<l.len(){
    if p.element==e{
      return i
    }
    p=p.next
    i++
  }
  return 0
}

func main() {
  x:=cll{}
	fmt.Println(x.len())
	fmt.Println(x.isempty())
	x.addlast(4)
	x.addlast(5)
	x.addlast(8)
	x.addfirst(2)
	x.addfirst(1)
	x.addany(3,3)
	x.addany(6,6)
	x.display()
	fmt.Println("\n deleted element is",x.removelast())
	fmt.Println("\n deleted element is",x.removefirst())
	x.display()
	fmt.Println("\n deleted element is",x.removeany(4))
	x.display()
	fmt.Println("\n 4 is found at index",x.search(4))
	fmt.Println("\n length is",x.len())
	
}
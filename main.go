package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

type LinkedList[T any] struct {
	next *List[T]
	len  int
}

func (l *LinkedList[T]) Add(aVal T) {
	temp := l.next

	for temp.next != nil {
		temp = temp.next
	}

	node := List[T]{}

	node.val = aVal
	temp.next = &node
}

//Function to print the list
func (l *LinkedList[T]) PrintList() {
	temp := l.next
	for temp.next != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
}

func (l *List[T]) Remove() {
	if l == nil {
		fmt.Println("List is empty")
	} else if l.next != nil {
		//remove node
		fmt.Println("Removed: ", l.val)
		l = nil
	} else {
		fmt.Println("Removed: ", l.val)
		fmt.Println("Deleted last item")
		l = nil
	}
}

func main() {
	l := LinkedList[string]{}
	l.Add("foo")
	l.Add("bar")
	l.Add("saving")
	l.Add("private")
	l.Add("ryan")
	fmt.Println("*****Actual List****")
	l.PrintList()
	fmt.Println("*********")
	//l.Remove()
	//l.PrintList()
}

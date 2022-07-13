package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Add(aVal T) {
	for l.next != nil {
		l = l.next
	}

	temp := List[T]{}

	if l.next == nil {
		l.val = aVal
		l.next = &temp
	}
}

//Function to print the list
func (l *List[T]) PrintList() {
	for l.next != nil {
		fmt.Println(l.val)
		l = l.next
	}
}

func main() {
	l := List[string]{}
	l.Add("foo")
	l.Add("bar")
	l.Add("saving")
	l.Add("private")
	l.Add("ryan")

	l.PrintList()
}

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
	node := List[T]{}
	node.val = aVal
	if l.next == nil {
		l.next = &node
	} else {
		temp := l.next

		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &node
	}

}

//Function to print the list
func (l *LinkedList[T]) PrintList() {
	if l.next == nil {
		return
	}
	temp := l.next
	for temp.next != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
	if temp.next == nil {
		fmt.Println(temp.val)
	}
}

func (l *LinkedList[T]) RemoveFirst() {
	if l.next == nil {
		fmt.Println("List is empty")
		return
	}

	temp := l.next
	fmt.Println("Removed: ", temp.val)

	if temp.next != nil {
		l.next = temp.next
	} else {
		l.next = nil
		fmt.Println("Deleted last item")
	}
	temp = nil
}

func main() {
	l := LinkedList[string]{}
	l.Add("foo")
	// l.Add("bar")
	// l.Add("saving")
	// l.Add("private")
	// l.Add("ryan")
	fmt.Println("*****Actual List****")
	l.PrintList()

	l.RemoveFirst()
	fmt.Println("*****After RemoveFirst****")
	l.PrintList()
	l.RemoveFirst()
	l.PrintList()
}

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

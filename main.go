package main

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Add(aVal T) {
	temp := List[T]{}

	if l.next == nil {
		l.val = aVal
		l.next = &temp
	} else {
		l.next.val = aVal
		l.next.next = &temp
	}

}

//Function to print the list
func (l *List[T]) PrintList() {

}

func main() {
	l := List[string]{}
	l.Add("foo")
	l.Add("bar")
	l.Add("savinf")
	l.Add("private")
	l.Add("ryan")
}

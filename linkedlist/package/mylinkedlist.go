package linkedlist

type LinkedList struct {
	Head *Node
}

type Node struct {
	Val  int
	Next *Node
}

func Init(val int) *LinkedList {
	ll := &LinkedList{}

	ll.Head = &Node{val, nil}
	return ll
}

func (ll *LinkedList) Append(val int) *LinkedList {
	newNode := &Node{val, nil}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	return ll
}

func (ll *LinkedList) GetFirst() int {
	return ll.Head.Val
}

func (ll *LinkedList) GetLast() int {
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Val
}

func (ll *LinkedList) GetLength() int {
	current := ll.Head
	length := 0
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

func (ll *LinkedList) Get(index int) int {
	current := ll.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Val
}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		print(current.Val, " --> ")
		current = current.Next
	}

	println("nil")
}

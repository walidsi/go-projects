package main

import linkedlist "github.com/walidsi/go-projects/linkedlist/package"

func main() {
	//var ll *linkedlist.List = linkedlist.Init(1)
	var ll *linkedlist.LinkedList = &linkedlist.LinkedList{Head: &linkedlist.Node{Next: nil, Val: 1}}
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)
	ll.Print()

	println("Length of the list: ", ll.GetLength())
	println("First element: ", ll.GetFirst())
	println("Last element: ", ll.GetLast())
	println("Element at index 2: ", ll.Get(2))
}

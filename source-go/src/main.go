package main

import (
	"algorithm/list"
	"fmt"
)

func main() {

	ll := list.NewLinkedList()
	ll.Push(11)
	fmt.Println(ll.Head.Next.Data)
	fmt.Println(ll.Length)

	ll.Push(22)
	fmt.Println(ll.Head.Next.Next.Data)
	fmt.Println(ll.Length)

	fmt.Println(ll.Insert(2, 66))
	fmt.Println(ll.Head.Next.Next.Data)
	fmt.Println(ll.Length)

	fmt.Println(ll.GetOneNode(2))
}

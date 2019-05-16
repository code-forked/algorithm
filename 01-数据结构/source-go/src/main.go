package main

import (
	"algorithm/list"
	"fmt"
)

func main() {
	sl := list.NewStaticList(5)
	fmt.Println(sl)

	ll := list.NewLinkedList()
	ll.Push(1)
	ll.Push(2)
	for i := 0; i < ll.Length; i++ {
		fmt.Println(ll.GetOneNode(i))
	}
}

package main

import (
	"fmt"
	"algorithm/array"
	"algorithm/list"
)


func main() {
	// 稀疏数组
	s1 := array.NewSparseArray(11,11)
	s2 := array.NewSparseArray(11,11)
	fmt.Printf("%p\n",&s1)
	fmt.Printf("%p\n",&s2)
	// line
	l1 := list.NewLineList(10)
	l2 := list.NewLineList(10)
	fmt.Printf("%p\n",&l1)
	fmt.Printf("%p\n",&l2)
}

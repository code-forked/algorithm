package main

import (
	"algorithm/array"
	SequenList "algorithm/list/SequenList"
	LinkedList "algorithm/list/LinkedList"
	"fmt"
)

func testSparseArray() {

	var primitArray [11][11]int
	primitArray[1][2] = 1
	primitArray[2][3] = 2
	array.ShowOrigin(primitArray)

	s := array.NewSparseArray(primitArray)
	array.ShowSparse(s)

	arr := array.TransToArray(s)
	array.ShowOrigin(arr)
}

func testSequenList() {

	sl := SequenList.New()
	sl.Show()

	sl.Append(7)
	sl.Append(9)
	sl.Show()

	sl.Insert(1,6)
	sl.Show()

	sl.Pop()
	sl.Show()

}

func testLinkedList() {

	ll := LinkedList.New()
	ll.Show()

	ll.Append(9)
	ll.Append(5)
	ll.Show()

	ll.Insert(1,7)
	ll.Insert(3,4)
	ll.Show()

	ll.Delete(3)
	ll.Show()
	
	fmt.Println(ll.Node(1))
}

func main() {

	fmt.Println("start run...")

	// testSparseArray()				// 测试稀疏数组

	// testSequenList()					// 测试顺序表	

	testLinkedList()					// 测试单链表

}

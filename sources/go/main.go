package main

import (
	"algorithm/list"
	"fmt"
)

func testSequenList() {

	sl := list.NewSequenList()
	sl.Show()

	sl.Append(7)
	sl.Append(9)
	sl.Show()

	sl.Insert(1,6)
	sl.Show()

	sl.Pop()
	sl.Show()

	// err := al.Insert(1,15)
	// fmt.Println(err)
	// al.Show()

	// err = al.Insert(2,13)
	// fmt.Println(err)
	// al.Show()
	// fmt.Println(al.Length())

}

func main() {

	fmt.Println("start run...")

	testSequenList()				// 测试顺序表	



}

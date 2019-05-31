package main

import (
	"algorithm/list"
	"fmt"
)


func main() {

	al := list.NewArrayList()
	al.Show()

	al.Append(10)
	al.Append(12)
	al.Show()

	err := al.Insert(1,15)
	fmt.Println(err)
	al.Show()

	err = al.Insert(2,13)
	fmt.Println(err)
	al.Show()
	fmt.Println(al.Length())



}

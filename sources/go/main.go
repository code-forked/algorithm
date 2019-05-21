package main

import (
	"fmt"
)

func Fbi(i int) int {
	if i == 0 || i == 1 {
		return i
	}
	return Fbi(i-1) + Fbi(i-2)
}

func main() {
	fmt.Println(Fbi(10))
}

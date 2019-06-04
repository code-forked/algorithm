package main

import (
	"algorithm/array"
	SequenList "algorithm/list/SequenList"
	LinkedList "algorithm/list/LinkedList"
	RingList "algorithm/list/RingList"
	DoublyList "algorithm/list/DoublyList"
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

func testRingList() {

	rl := RingList.New()
	rl.Show()

	rl.Append(9)
	rl.Append(5)
	rl.Show()

	rl.Insert(1,7)
	rl.Insert(3,4)
	rl.Show()

	rl.Delete(3)
	rl.Show()
	
}

func testDoublyList() {

	dl := DoublyList.New()
	dl.Show()

	dl.Append(7)
	dl.Append(3)
	dl.Append(9)
	dl.Show()

	dl.Insert(1,5)
	dl.Insert(3,6)
	dl.Show()

	fmt.Println(dl.Delete(3))
	dl.Show()

	fmt.Println(dl.Delete(4))		// 删除最后一个节点
	dl.Show()
}

// 约瑟夫环
func testJosephus(){	

	// 生成循环链表
	rl := RingList.New()
	for i := 0; i < 41; i++ {					// 多个人组成的循环链表
		rl.Append(i + 1)						// 每个人都对应自己的原始编号
	}
	fmt.Print("生成的循环链表为：")
	rl.Show()

	// 每次读取3个元素，读取位置从startNode开始
	startNode := rl.GetHead()				// 起始开始报数的元素
	delIndex := 0 								// 要删除元素的索引

	for {

		for i:= 1; i <= 3; i++ {	

			if rl.GetNext(startNode) == rl.GetHead() {
				delIndex = 1								// 每次从头结点开始时，删除索引变为0，重新记录
				startNode = rl.GetNext(rl.GetNext(startNode))
			} else {
				startNode = rl.GetNext(startNode)
				delIndex ++	
			}
			
		}

		// 删除该元素
		val, _ := rl.Node(delIndex)
		fmt.Println("删除数据为：", delIndex, " ", val)
		rl.Delete(delIndex)

		// 如果只有一个元素
		if rl.Length() < 3 {
			break;
		}
	}

	fmt.Print("最终的循环链表为：")
	rl.Show()
	return
}

func main() {

	fmt.Println("start run...")

	// testSparseArray()				// 测试稀疏数组

	// testSequenList()					// 测试顺序表	

	// testLinkedList()					// 测试单链表

	// testRingList()					// 测试循环链表

	testJosephus()						// 测试约瑟夫环

	// testDoublyList()					// 测试双向链表

}

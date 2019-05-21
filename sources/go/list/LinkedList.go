/**
 * 单链表
 */

package list

import "errors"

// 先设计一个结点对象，存储单链表上某个结点数据
type node struct {
	Data interface{}			// 数据域
	Next *node					// 指针域
}

// 再设计单链表对象
type LinkedList struct {
	Head *node					// 单链表头，有了头结点，就能找到所有结点
	Length int					// 当前链表中数据元素数量
}

// 创建单链表
func NewLinkedList() *LinkedList {
	head := &node{0, nil}
	return &LinkedList{
		head,
		0,
	}
}

// 判断空
func (ll *LinkedList)IsEmpty() bool {
	return ll.Head.Next == nil
}

// 增加：从尾部增加一个结点
func (ll *LinkedList) Push(o interface{}){
	appendNode := &node{o, nil}			// 要插入的结点
	lastNode := ll.Head.Next						// 查找最后一个结点
	if lastNode == nil {							// 第一次添加
		ll.Head.Next = appendNode
		ll.Length ++
		return
	}
	for lastNode.Next != nil {						// 不是第一次添加
		lastNode = lastNode.Next
	}

	lastNode.Next = appendNode
	ll.Length ++
}

// 增加：任意位置插入结点
func (ll *LinkedList) Insert(index int, o interface{}) (bool, error){
	if index < 0 || index > ll.Length {
		return false, errors.New("无效的插入位置")
	}
	currentNode := ll.Head
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next		// 找到要插入的位置
	}
	appendNode := &node{Data:o, Next:nil}
	appendNode.Next = currentNode.Next
	currentNode.Next = appendNode
	ll.Length ++
	return true, nil

}

// 删除：删除指定位置结点
func (ll *LinkedList) Delete(index int) (bool, error) {
	if index < 0 || index >= ll.Length {
		return false, errors.New("指定索引位置有误")
	}
	currentNode := ll.Head
	var beforeNode *node
	for i := 0; i <= index; i++ {
		beforeNode, currentNode = currentNode,  currentNode.Next
	}
	beforeNode.Next = currentNode.Next
	currentNode = nil
	ll.Length --
	return true, nil
}

// 查询： 获取指定位置结点
func (ll *LinkedList) GetOneNode(index int) (*node, error) {
	if index < 0 || index >= ll.Length {
		return nil, errors.New("索引越界")
	}
	currentNode := ll.Head
	for i := 0; i <= index; i++ {
		currentNode = currentNode.Next
	}
	return currentNode, nil
}

// 清空链表
func (ll *LinkedList) Clear() {
	currentNode := ll.Head.Next
	for currentNode != nil {
		temp := currentNode.Next
		currentNode = nil
		currentNode = temp
	}
	ll.Head.Next = nil
}

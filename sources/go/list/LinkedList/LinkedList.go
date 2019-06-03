/**
 * 单链表
 */

package list

import "errors"

// 先设计一个结点对象，存储单链表上某个节点数据
type node struct {
	data interface{}			// 数据域
	next *node					// 指针域
}

// 单链表对象
type LinkedList struct {
	head *node					// 单链表头，有了头节点，就能找到所有结点
	length int					// 当前链表中数据元素数量
}

// 创建单链表
func New() *LinkedList {
	head := &node{0, nil}
	return &LinkedList{
		head,
		0,
	}
}

// 判断空
func (ll *LinkedList)IsEmpty() bool {
	return ll.head.next == nil
}

// 增加：末尾添加
func (ll *LinkedList) Append(data interface{}){

	insertNode := &node{data, nil}					// 要插入的节点

	// 查询最后一个节点
	lastNode := ll.head.next
	if lastNode == nil {							// 第一次添加
		ll.head.next = insertNode
		ll.length ++
		return
	}
	for lastNode.next != nil {						// 不是第一次添加
		lastNode = lastNode.next
	}
	lastNode.next = insertNode
	ll.length ++
	return
}

// 增加：任意位置插入结点
func (ll *LinkedList) Insert(index int, data interface{}) error{

	if index < 0 || index > ll.length {
		return errors.New("index overflow")
	}

	currentNode := ll.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next		// 找到要插入的位置
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

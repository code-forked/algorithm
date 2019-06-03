/**
 * 双向链表
 */

 package list

 import (
	//  "errors"
	 "fmt"
 )
 
 // 节点对象
 type node struct {
	 data interface{}
	 prev *node			
	 next *node					
 }
 
 // 单链表对象：存储头节点即可
 type DoublyList struct {
	 head *node					// 双向链表表头，笔者这里表头用来存储个数
	 tail *node					// 双向链表表尾
 }
 
 // 创建单链表
 func New() *DoublyList {
	 head := &node{0, nil, nil,}		// 头节点存储链表中元素的个数
	 tail := &node{nil, head, nil}
	 head.next = tail
	 return &DoublyList{
		 head,
		 tail,
	 }
 }
 
 // 增加：末尾添加
 func (dl *DoublyList) Append(data interface{}){

	var len int = 0
	len = dl.head.data.(int)

	insertNode := &node{data, dl.tail, nil,}					// 要插入的节点
	dl.tail = insertNode

	len++
	dl.head.data = len

	 return
 }
 
//  // 增加：任意位置插入结点
//  func (ll *RingList) Insert(index int, data interface{}) error{
 
// 	 var len int = 0
// 	 len = ll.head.data.(int)
 
// 	 if index < 1 || index > len {
// 		 return errors.New("index overflow")
// 	 }
 
// 	 beforeNode := ll.head
// 	 appendNode := &node{data, nil}
 
// 	 for i := 0; i < index - 1; i++ {
// 		 beforeNode = beforeNode.next		// 找到要插入的位置的前置元素
// 	 }
 
// 	 appendNode.next = beforeNode.next
// 	 beforeNode.next = appendNode
 
// 	 len ++
// 	 ll.head.data = len
 
// 	 return nil
 
//  }
 
//  // 删除：删除指定位置结点
//  func (ll *RingList) Delete(index int) (interface{}, error) {
 
// 	 var len int = 0
// 	 len = ll.head.data.(int)
 
// 	 if index < 0 || index >= len {
// 		 return nil,errors.New("index overflow")
// 	 }
 
// 	 currentNode := ll.head
// 	 var beforeNode *node
// 	 var delData interface{}					// 被删除的数据
 
// 	 for i := 0; i < index; i++ {
// 		 beforeNode = currentNode
// 		 currentNode = currentNode.next
// 	 }
 
// 	 beforeNode.next = currentNode.next
// 	 currentNode = nil
 
// 	 len--
// 	 ll.head.data = len
 
// 	 return delData, nil
//  }
 
 // 打印链表
 func (dl *DoublyList) Show() {

	 var len int = 0
	 len = dl.head.data.(int)

	 currentNode := dl.head

	 for i := 0; i <= len-1; i++ {
		 fmt.Print(currentNode.data)
		 if i == len - 1 {
			 fmt.Print(" \n")
		 } else {
			 fmt.Print(" ")
		 }
 
		 currentNode = currentNode.next
	 }
 }
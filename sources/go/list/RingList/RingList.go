/**
 * 循环链表
 */

 package list

// 循环链表节点
 type node struct {				
	data interface{}
	next *node					
 }

 // 循环链表
 type RingList struct {
	head *node						// 单链表头，有了头结点，就能找到所有结点
	length int 						// 循环链表长度
 }
 
 // 初始化线循环链表
 func New() *RingList {
	return &RingList{
		nil,
		0,
	}
 }

 // 插入节点
func (rl *RingList)Append(data interface{}){

	n := &node{
		data,
		nil,
	}

	if rl.Length() == 0 {
		rl.
	}

}

// 删除节点

// 获取节点

 // 判断链表是否为空
 func (rs *RingList)IsEmpty() bool {
	 if rs.length == 0 {
		 return true
	 }
	 return false
 }

 // 获取链表长度
 func (rs *RingList)Length() int {
	return rs.length
 }

// 遍历链表
func (rl *RingList)Show() {

}
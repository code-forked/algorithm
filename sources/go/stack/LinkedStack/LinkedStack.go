/**
 * 栈:链栈
 */

package LinkedStack

type Node struct {
	data interface{}
	top *Node			// 栈顶指针
}

type LinkedStack struct {
	Top    *Node // 栈顶元素
	Length int
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		nil,
		0,
	}
}

// 压栈
func (ls *LinkedStack) Push(data interface{}) {
	node := new(Node)
	node.data = data
	node.next = ls.Top
	ls.Top = node
	ls.Length++
}

// 出栈
func (ls *LinkedStack) Pop() interface{} {
	if ls.Length == 0 {
		panic("当前栈是空栈")
	}
	value := ls.Top.data
	node := ls.Top
	ls.Top = node.next
	ls.Length--
	return value
}

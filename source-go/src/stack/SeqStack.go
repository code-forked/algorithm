/**
 * 栈:顺序结构
 */

package stack

type SeqStack struct {
	Data   []interface{} // 线性表内数据
	Length int
}

func NewSeqStack() *SeqStack {
	s := make([]interface{}, 0)
	return &SeqStack{
		s,
		0,
	}
}

// 压栈
func (ss *SeqStack) Push(data interface{}) {
	ss.Data = append(ss.Data, data)
	ss.Length++
}

// 出栈
func (ss *SeqStack) Pop() interface{} {
	if ss.Length == 0 {
		panic("当前栈是空栈")
	}
	index := ss.Length - 1
	value := ss.Data[index]
	ss.Data = append(ss.Data[:index])
	ss.Length--
	return value
}

// 清空栈
func (ss *SeqStack) Clear() {
	ss = NewSeqStack()
}

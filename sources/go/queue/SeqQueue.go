package queue

// 顺序结构存储队列
type SeqQueue struct {
	Size int
	Length int
	Data []interface{}
	Front int 					// 表示指向队列队首
	Rear int 					// 表示指向队列尾部
}

func NewSeQueue(size int) {
	return &SeqQueue{
		size,

	}
}

// 向队列添加数据
package array

type Node struct {
	row int
	col int
	val interface{}
}

func NewSparseArray(row int, col int) []Node{
	var sa []Node
	n := Node{
		row: row,
		col: col,
		val: nil,
	}
	sa = append(sa, n)
	return sa
}

// 普通数组转换稀疏数组


// 稀疏数组转普通数组
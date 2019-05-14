/**
 * 线性表
 */

package list

import (
	"errors"
	"fmt"
)

// 线性表结构体对象
type LineList struct {
	MaxSize int					// 该线性表最大容量
	Length int					// 该线性表当前长度
	Data []interface{}			// 线性表内数据
}

// 初始化线性表函数
func NewLineList(size int) *LineList {
	return &LineList{
		MaxSize: size,
		Length: 0,
		Data: make([]interface{}, size),
	}
}

// 判断线性表是否为空方法
func (l *LineList)IsEmpty() bool {
	if l.Length == 0 {
		return true
	}
	return false
}

// 判断线性表是否已满
func (l *LineList)IsFull() bool {
	if l.Length == l.MaxSize {
		return true
	}
	return false
}

// 判断线性表索引是否越界
func (l *LineList) isOver(index int) bool {
	if index < 1 || index > l.Length {
		return true
	}
	return false
}

// 插入元素：从末尾append一个数据
func (l *LineList)Append(data interface{})(bool, error) {
	if ok := l.IsFull(); ok {
		return false, errors.New("线性表已满，无法添加数据")
	}
	l.Data = append(l.Data, data)
	l.Length ++
	return true, nil
}

// 插入元素：任意位置插入元素
func (l *LineList)Insert(index int, data interface{}) (bool, error) {
	if ok := l.IsFull(); ok {
		return false, errors.New("线性表已满，无法添加数据")
	}
	if ok := l.isOver(index); ok {
		return false, errors.New("插入索引越界")
	}
	b,_ := l.Append("")			// 增加一个空数据，防止下面访问越界
	fmt.Println("插入临时空数据结果:", b)
	for j := l.Length - 1; j > index - 1; j -- {
		// 从后往前赋值，新增一个空节点，然后把数据一个个后移，直到插入位置
		l.Data[j] = l.Data[ j - 1]
	}
	l.Data[index - 1] = data
	return true, nil
}

// 删除元素：从末尾pop一个数据
func (l *LineList)Pop() (interface{}, error) {
	if ok := l.IsEmpty(); ok {
		return "", errors.New("表中没有任何元素")
	}
	lastE := l.Data[l.Length - 1]
	l.Data = l.Data[:l.Length - 1]
	l.Length --
	return lastE, nil
}

// 删除元素：从任意位置删除一个数据
func (l *LineList)Delete(index int) (interface{}, error) {
	if ok := l.isOver(index); ok {
		return "", errors.New("索引不在线性表范围内")
	}
	if ok := l.IsEmpty(); ok {
		return "", errors.New("空表没有课删除的数据")
	}
	deleteE := l.Data[index - 1]
	for j := index - 1; j < l.Length - 1; j++ {
		l.Data[j] = l.Data[ j + 1]
	}
	l.Data = l.Data[:l.Length - 1]
	l.Length --
	return deleteE,  nil
}

// 修改元素

// 查询某个元素
func (l *LineList)GetByIndex(index int) (interface{}, error) {
	if ok := l.isOver(index); ok {
		return "", errors.New("查询索引越界")
	}
	return l.Data[index - 1], nil
}



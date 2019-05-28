/**
 * 线性表:顺序结构存储
 */

package list

import (
	"errors"
	"fmt"
)

// 线性表结构体对象
type SequenList struct {
	maxSize int									// 该线性表最大容量
	length int									// 该线性表当前长度
	data []interface{}							// 线性表内数据
}

// 初始化线性表函数
func NewSequenList(size int) *SequenList {
	return &SequenList{
		maxSize: size,
		length: 0,
		data: make([]interface{}, 0),
	}
}

// 打印线性表
func (sl *SequenList)Show() {
	fmt.Println(sl)
}

// 插入元素：从末尾append一个数据
func (sl *SequenList)Append(data interface{})  error{
	if sl.IsFull() {
		return errors.New("SequenList overflow")
	}
	sl.data = append(sl.data, data)
	sl.length ++
	return nil
}

// 插入元素：任意位置插入元素
func (sl *SequenList)Insert(index int, data interface{}) error {
	if sl.IsFull() {
		return errors.New("SequenList overflow")
	}
	if sl.isOver(index) {
		return errors.New("index overflow")
	}

	b,_ := sl.Append("")								// 增加一个空数据，防止下面访问越界
	for i := 0; i < sl.length - 1; i {
		// 从后往前赋值，新增一个空节点，然后把数据一个个后移，直到插入位置
		sl.data[j] = sl.data[ j - 1]
	}
	sl.data[index - 1] = data
	return true, nil
}

// 删除元素：从末尾pop一个数据
func (l *SequenList)Pop() (interface{}, error) {
	if ok := l.IsEmpty(); ok {
		return "", errors.New("表中没有任何元素")
	}
	lastE := l.data[l.length - 1]
	l.data = l.data[:l.length - 1]
	l.length --
	return lastE, nil
}

// 删除元素：从任意位置删除一个数据
func (l *SequenList)Delete(index int) (interface{}, error) {
	if ok := l.isOver(index); ok {
		return "", errors.New("索引不在线性表范围内")
	}
	if ok := l.IsEmpty(); ok {
		return "", errors.New("空表没有课删除的数据")
	}
	deleteE := l.data[index - 1]
	for j := index - 1; j < l.length - 1; j++ {
		l.data[j] = l.data[ j + 1]
	}
	l.data = l.data[:l.length - 1]
	l.length --
	return deleteE,  nil
}

// 修改元素

// 查询某个元素
func (l *SequenList)GetByIndex(index int) (interface{}, error) {
	if ok := l.isOver(index); ok {
		return "", errors.New("查询索引越界")
	}
	return l.data[index - 1], nil
}





// 判断线性表是否为空方法
func (sl *SequenList)IsEmpty() bool {
	if sl.length == 0 {
		return true
	}
	return false
}

// 判断线性表是否已满
func (sl *SequenList)IsFull() bool {
	if sl.length == sl.maxSize {
		return true
	}
	return false
}

// 判断线性表索引是否越界
func (sl *SequenList)isOver(index int) bool {
	if index < 1 || index > sl.length {
		return true
	}
	return false
}

// 获取线性表长度
func (sl *SequenList)Length() int{
	return sl.length
}

// 获取线性表容量
func (sl *SequenList) Size() int{
	return sl.maxSize
}

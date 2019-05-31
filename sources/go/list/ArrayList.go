/**
 * 线性表:顺序结构存储
 */

package list

import (
	"errors"
	"fmt"
)

var arr [5]int

// 线性表结构体对象
type ArrayList struct {
	size int									// 该线性表最大容量
	length int									// 该线性表最大长度
	data [5]int								// 线性表内数据，这里为了演示默认设置为100长度的int数组，所有元素默认为0(Go的0值机制)
}

// 初始化线性表函数
func NewArrayList() *ArrayList {						
	return &ArrayList{
		size: 5,
		length: 0,
		data: arr,
	}
}

// 打印线性表
func (sl *ArrayList)Show() {
	fmt.Println(sl)
}

// 插入元素：从末尾append一个数据
func (al *ArrayList)Append(data int)  error{

	// 判断空间是否已满
	if al.IsFull() {
		return errors.New("SequenList overflow")
	}

	// 从末尾插入：由于数组中所有元素已经默认是0了，那么第一个不是0的元素就是最后一位
	for i := 0; i <= al.size; i++ {
		if al.data[i] == 0 {
			al.data[i] = data
			al.length++
			break
		}
	}

	return nil
}

// 插入元素：任意位置插入元素
func (al *ArrayList)Insert(index int, data int) error {

	if al.IsFull() {
		return errors.New("SequenList overflow")
	}

	if al.isOver(index) {
		return errors.New("index overflow")
	}

	// 情况一：在最后一位插入
	if index == al.length - 1{
		al.data[index] = data
		al.length++
		return nil
	}

	// 情况二：在其他位置插入
	currentItem := al.data[index]						
	nextItem := al.data[index+1]
	al.data[index] = data							// 插入数据
	al.length++

	for i := index + 1; i <= al.length; i++ {

		fmt.Println("i===",i)
		fmt.Println("值===",al.data[i])

		al.data[i] = currentItem

		if i == al.length {							// 循环到最后一位
			break	
		}
		
		currentItem = nextItem
	
		tmp := al.data[i + 1]
		al.data[i + 1] = nextItem					// 后面的数据全部后移
		nextItem = tmp

	}

	return nil
}

// 删除元素：从末尾pop一个数据
func (al *ArrayList)Pop() (int, error) {

	if al.IsEmpty() {
		return 0, errors.New("object is empty")
	}

	e := al.data[al.length - 1]
	al.data[al.length - 1] = 0
	al.length --
	return e, nil
}

// // 删除元素：从任意位置删除一个数据
// func (l *SequenList)Delete(index int) (interface{}, error) {
// 	if ok := l.isOver(index); ok {
// 		return "", errors.New("索引不在线性表范围内")
// 	}
// 	if ok := l.IsEmpty(); ok {
// 		return "", errors.New("空表没有课删除的数据")
// 	}
// 	deleteE := l.data[index - 1]
// 	for j := index - 1; j < l.length - 1; j++ {
// 		l.data[j] = l.data[ j + 1]
// 	}
// 	l.data = l.data[:l.length - 1]
// 	l.length --
// 	return deleteE,  nil
// }

// // 修改元素

// // 查询某个元素
// func (l *ArrayList)GetByIndex(index int) (interface{}, error) {
// 	if ok := l.isOver(index); ok {
// 		return "", errors.New("查询索引越界")
// 	}
// 	return l.data[index - 1], nil
// }







// 下列为常用工具方法

// 获取线性表长度
func (al *ArrayList)Length() int{
	return al.length
}

// 判断线性表是否已满
func (sl *ArrayList)IsFull() bool {
	if sl.length == sl.size {
		return true
	}
	return false
}

// 判断线性表是否为空方法
func (al *ArrayList)IsEmpty() bool {
	if al.length == 0 {
		return true
	}
	return false
}

// 获取线性表容量
func (sl *ArrayList)Size() int{
	return sl.size
}

// 判断线性表索引是否越界
func (al *ArrayList)isOver(index int) bool {
	if index < 0 || index > al.length - 1 {
		return true
	}
	return false
}

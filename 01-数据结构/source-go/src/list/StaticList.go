/**
 * 静态链表
 */

package list

import (
	"errors"
	"fmt"
	"os"
)

// 结点结构体
type component struct {
	Data interface{}
	Cur  int // 游标：为0时表示无指向
}

// 静态链表
type StaticList struct {
	Arr     []component
	MaxSize int
	Length  int
}

func NewStaticList(size int) *StaticList {

	if size < 3 {
		fmt.Println("参数错误")
		return nil
	}

	list := make([]component, size)

	for i := 0; i < size-1; i++ {
		list[i].Cur = i + 1
	}
	list[size-2].Cur = 0
	list[size-1].Cur = 0 // 初始化时，静态链表为空，最后一个元素cur为0
	return &StaticList{list, size, 0}
}

// 判断是否为空
func (sl *StaticList) IsEmpty() bool {
	if sl.Length == 0 {
		return true
	}
	return false
}

// 分配结点
func (sl *StaticList) malloc() int {
	i := sl.Arr[0].Cur
	if i == 0 {
		os.Exit(0)
	}
	sl.Arr[0].Cur = sl.Arr[i].Cur
	return i
}

// 回收结点
func (sl *StaticList) free(index int) {
	sl.Arr[index].Cur = sl.Arr[0].Cur
	sl.Arr[0].Cur = index
}

// 回收链表到备用链表
func (sl *StaticList) DestroyList() {
	if sl.Arr[sl.MaxSize-1].Cur == 0 {
		return
	}
	j := sl.Arr[sl.MaxSize-1].Cur
	sl.Arr[sl.MaxSize-1].Cur = 0
	i := sl.Arr[0].Cur
	sl.Arr[0].Cur = j
	if j > 0 {
		j = sl.Arr[j].Cur
	}
	sl.Arr[j].Cur = i
}

// 插入节点
func (sl *StaticList) Insert(data interface{}, index int) (bool, error) {
	if index < 1 || index > sl.Length {
		return false, errors.New("插入索引不合法")
	}
	i := sl.Arr[sl.MaxSize-1].Cur
	j := 1
	for i > 0 && j < index-1 {
		j++
		i = sl.Arr[i].Cur
	}
	tmp := sl.Arr[i].Cur
	cur := sl.malloc()
	sl.Arr[cur].Data = data
	sl.Arr[cur].Cur = tmp
	sl.Arr[i].Cur = cur
	return true, nil
}

// 显示链表结构
func (sl *StaticList) Traverse() {
	for i, v := range sl.Arr {
		fmt.Printf("%5d:%5d,%5s", i, v.Cur, v.Data)
	}
}

// 获取数据元素位置
func (sl *StaticList) GetLocation(data interface{}) int {
	location := 0
	i := sl.Arr[sl.MaxSize-1].Cur
	for i > 0 {
		location++
		if sl.Arr[i].Data == data {
			return location
		}
		i = sl.Arr[i].Cur
	}
	return location
}

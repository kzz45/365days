package arraylist

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	Size() int                                   // 数组大小
	Get(index int) (interface{}, error)          // 获取第几个元素
	Set(index int, new_val interface{}) error    // 修改数据
	Insert(index int, new_val interface{}) error // 插入数据
	Append(new_val interface{})                  // 追加
	Clear()                                      // 清空
	Delete(index int) error                      // 删除
	String() string                              //返回字符串
}

// 数据结构 字符串 整数
type ArrayList struct {
	dataStore []interface{} // 数组存储
	TheSize   int           // 数组大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      // 初始化结构体
	list.dataStore = make([]interface{}, 0, 10) // 开辟空间10个
	list.TheSize = 0
	return list
}

func (list *ArrayList) check_is_full() {
	if list.TheSize == cap(list.dataStore) {
		new_dataStore := make([]interface{}, 2*list.TheSize) // 开辟双倍内存
		copy(new_dataStore, list.dataStore)
		list.dataStore = new_dataStore
	}
}

func (list *ArrayList) Size() int {
	return list.TheSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Append(new_val interface{}) {
	list.dataStore = append(list.dataStore, new_val)
	list.TheSize++

}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore...)
}

func (list *ArrayList) Set(index int, new_val interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = new_val

	return nil
}

func (list *ArrayList) Insert(index int, new_val interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.check_is_full()                             // 检测内存 如果满了自动追加
	list.dataStore = list.dataStore[:list.TheSize+1] // 插入数据时 内存移动一位
	for i := list.TheSize; i > index; i-- {          // 从后往前移动
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = new_val // 插入数据
	list.TheSize++                  // 索引追加
	return nil
}

func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.TheSize--
	return nil
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10) // 开辟空间
	list.TheSize = 0
}

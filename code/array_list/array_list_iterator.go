package arraylist

import "errors"

type Iterator interface {
	HasNext() bool              // 是否有下一个
	Next() (interface{}, error) // 下一个
	Remove()                    //删除
	GetIndex() int              // 得到索引
}

type Iterable interface {
	Iterator() Iterator // 构造初始化接口
}

// 构造指针访问数组
type ArrayListIterator struct {
	list          *ArrayList // 数组指针
	current_index int        // 当前索引
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator) // 构造迭代器
	it.current_index = 0
	it.list = list
	return it
}

func (it *ArrayListIterator) HasNext() bool {
	return it.current_index < it.list.TheSize
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个")
	}
	val, err := it.list.Get(it.current_index)
	it.current_index++
	return val, err
}

func (it *ArrayListIterator) Remove() {
	it.current_index--
	it.list.Delete(it.current_index)
}

func (it *ArrayListIterator) GetIndex() int {
	return it.current_index
}

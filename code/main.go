package main

import (
	arraylist "365days/code/array_list"
	"fmt"
)

// func main1() {
// 	list := arraylist.NewArrayList()
// 	list.Append(1)
// 	list.Append(2)
// 	list.Append(3)
// 	fmt.Println(list, list.TheSize)
// }

// func main2() {
// 	list := arraylist.NewArrayList()
// 	list.Append("a")
// 	list.Append("b")
// 	list.Append("c")
// 	fmt.Println(list, list.TheSize)
// }

// func main3() {
// 	// 定义接口对象 赋值的对象必须实现接口的所有方法
// 	var list arraylist.List = arraylist.NewArrayList()
// 	list.Append("a")
// 	list.Append("b")
// 	list.Append("c")
// 	fmt.Println(list)
// }

// func main4() {
// 	// 定义接口对象 赋值的对象必须实现接口的所有方法
// 	var list arraylist.List = arraylist.NewArrayList()
// 	list.Append("a")
// 	list.Append("b")
// 	list.Append("c")
// 	fmt.Println(list)
// 	list.Insert(1, "d")
// 	fmt.Println(list)
// }

func main() {
	// 定义接口对象 赋值的对象必须实现接口的所有方法
	var list arraylist.List = arraylist.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	for i := 0; i < 100; i++ {
		list.Insert(1, "F")
		fmt.Println(list)
	}
	list.Insert(1, "D")
	fmt.Println(list)
}

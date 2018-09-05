package main

import (
	"fmt"
)

func main() {

	//数组可以不显示指定数组长度, 而用...代替
	arr1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	//数组的奇怪初始化方法
	arr2 := [...]int{4: -1} //它表示数组的前四个元素为默认值，最后一个为-1，共5个元素
	fmt.Println(arr2)

	x := []int{2, 3, 5, 7, 11}
	fmt.Println(x)
	// 数组的slice并不会实际复制一份数据，它只是创建一个新的数据结构，包含了另外的一个指针，一个长度和一个容量数据
	//[low:high] 用数学集合的方式来讲，就是[low, high)，即左闭右开。
	y := x[1:2] // x[1:2]   [3]      x[1:3]   [3,5]
	fmt.Println(y)
	fmt.Println(y[0])
	y[0] = 9
	fmt.Println(y[0])
	fmt.Println(y)
	fmt.Println(x)
	z := append(y, 4, 5)
	fmt.Println(z)
	fmt.Println(len(z))
	//
	var theSlice []int            //声明
	theSlice = make([]int, 5, 10) //开辟空间并附默认值
	for i, item := range theSlice {
		fmt.Printf("theSlice[%d]=%d\n", i, item)
	}
	//预留的空间需要重新切片才可以使用
	a := make([]int, 10, 20)
	fmt.Printf("%d, %d\n", len(a), cap(a))
	fmt.Println(a)
	b := a[:cap(a)]
	fmt.Println(b)
	a[0] = 1
	fmt.Println(a) //a[0]=1
	fmt.Println(b) //b[0]=1

	slice1 := []int{1, 2, 3} //声明并赋值
	slice2 := make([]int, 2) //声明并开辟空间
	slice3 := make([]int, 4) //声明并开辟空间

	copy(slice2, slice1)
	copy(slice3, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	sliceByte := append([]byte("hello "), "world"...)
	fmt.Println(sliceByte)
}

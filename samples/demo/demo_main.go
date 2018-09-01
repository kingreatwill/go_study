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
}

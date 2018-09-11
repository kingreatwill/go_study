package main

import (
	"unsafe"
	"fmt"
)
// 在Golang中，[]int 和 []MyInt是两种不同的类型。
// 因此，[]int的值不能转换为[]MyInt，反之亦然。 但是在unsafe.Pointer的帮助下，转换是可能的：
func main() {
	type MyInt int

	a := []MyInt{0, 1, 2}
	// b := ([]int)(a) // error: cannot convert a (type []MyInt) to type []int
	b := *(*[]int)(unsafe.Pointer(&a))

	b[0] = 3

	fmt.Println("a =", a) // a = [3 1 2]
	fmt.Println("b =", b) // b = [3 1 2]

	a[2] = 9

	fmt.Println("a =", a) // a = [3 1 9]
	fmt.Println("b =", b) // b = [3 1 9]
}
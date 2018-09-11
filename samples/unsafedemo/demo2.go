package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var x struct {
		a bool
		b int16
		c []int
	}

	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"
	/*
	上面的写法尽管很繁琐，但在这里并不是一件坏事，因为这些功能应该很谨慎地使用。
	不要试图引入一个uintptr类型的临时变量，因为它可能会破坏代码的安全性（译注：这是真正可以体会unsafe包为何不安全的例子）。
	下面段代码是错误的：test2()
	*/
	test2()

	test3()
}

func test2() {

	var x struct {
		a bool
		b int16
		c []int
	}

	// NOTE: subtly incorrect!
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 42
	fmt.Println(x.b) // "42"
}
//在这个例子中的转换可能是无意义的，但它是安全和合法的。
func test3() {
	var n int64 = 5
	var pn = &n
	var pf = (*float64)(unsafe.Pointer(pn))
	fmt.Println(*pf) //2.5e-323
	*pf = 3.1415
	fmt.Println(n) //4614256447914709615
}
//uintptr 和 unsafe.Pointer 的互相转换，
func test4() {
	a := [4]int{0, 1, 2, 3}
	p := &a[1] // 内存地址
	p1 := unsafe.Pointer(p)
	p2 := uintptr(p1)
	p3 := unsafe.Pointer(p2)
	fmt.Println(p1) // 0xc420014208
	fmt.Println(p2) // 842350543368
	fmt.Println(p3) // 0xc420014208
}
package main

import "fmt"

/*
new 和 make 都可以用来分配空间，初始化类型，但是它们确有不同。
new(T) 返回的是 T 的指针
new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，返回的是新值的地址，也就是 T 类型的指针 *T，该指针指向 T 的新分配的零值。


make 只能用于 slice，map，channel 三种类型
make(T, args) 返回的是初始化之后的 T 类型的值，这个新值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。
*/
func main() {
	//1
	p1 := new(int)
	fmt.Printf("p1 --> %#v \n ", p1)           //(*int)(0xc42000e250)
	fmt.Printf("p1 point to --> %#v \n ", *p1) //0
	// 等价1
	var p2 *int
	i := 0
	p2 = &i
	fmt.Printf("p2 --> %#v \n ", p2)           //(*int)(0xc42000e278)
	fmt.Printf("p2 point to --> %#v \n ", *p2) //0

	*p1 = 5
	fmt.Printf("p1  %#v \n ", *p1)
}

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)
/*
ArbitraryType 是int的一个别名，但是golang中，对ArbitraryType赋予了特殊的意义，千万不要死磕这个后边的int类型。
通常的说，我们把interface{}看作是任意类型，是比较浪荡的型号，具有老少通吃的特点，那么ArbitraryType这个类型，在golang系统中，是人兽皆不在话下的类型。比interface{}还要随意。

Pointer 是int指针类型的一个别名，在golang系统中，可以把Pointer类型，理解成任何指针的亲爹。

看这两个变量，没什么特殊的。提示一下：golang的指针类型长度与int类型长度，在内存中占用的字节数是一样的哟。
提示：
ArbitraryType类型的变量也可以是指针。所以，千万不要死磕type后边的那个int。

https://studygolang.com/articles/9446
*/
func main(){
	s :="123csdfsdfsd代诗歌是的dgsdgg13131"

	//fmt.Println(s)
	//Sizeof返回类型v本身数据所占用的字节数。返回值是“顶层”的数据占有的字节数。例如，若v是一个切片，它会返回该切片描述符的大小，而非该切片底层引用的内存的大小。
	fmt.Println(unsafe.Sizeof(s))
	//Alignof返回类型v的对齐方式（即类型v在内存中占用的字节数）；若是结构体类型的字段的形式，它会返回字段f在该结构体中的对齐方式。
	fmt.Println(unsafe.Alignof(s))
	//fmt.Println(unsafe.Offsetof(s))

	var x struct {
		a bool
		b int16
		c []int
	}

	//通常情况下布尔和数字类型需要对齐到它们本身的大小(最多8个字节),其它的类型对齐到机器字大小.(64位的机器字大小为64位,8字节)

	fmt.Printf("%-30s%-30s%-30s%-50s\n",
		"Row", "Sizeof", "Alignof(对齐倍数)", "Offsetof(偏移量)")

	fmt.Printf("%-30s%-30d%-30d%-50s\n",
		"x", unsafe.Sizeof(x), unsafe.Alignof(x), "")
	fmt.Printf("%-30s%-30d%-30d%-50d\n",
		"x.a", unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Printf("%-30s%-30d%-30d%-50d\n",
		"x.b", unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Printf("%-30s%-30d%-30d%-50d\n",
		"x.c", unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))


	fmt.Printf("%#016x\n", Float64bits(1.0)) // "0x3ff0000000000000"

}
func Float64bits(f float64) uint64 {
	fmt.Println(reflect.TypeOf(unsafe.Pointer(&f)))  //unsafe.Pointer
	fmt.Println(reflect.TypeOf((*uint64)(unsafe.Pointer(&f))))  //*uint64
	return *(*uint64)(unsafe.Pointer(&f))
}

package main

import "fmt"

/*
new 和 make 都可以用来分配空间，初始化类型，但是它们确有不同。
new(T) 返回的是 T 的指针
new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，返回的是新值的地址，也就是 T 类型的指针 *T，该指针指向 T 的新分配的零值。


make 只能用于 slice，map，channel 三种类型
make(T, args) 返回的是初始化之后的 T 类型的值，这个新值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。

Slice: 第二个参数 size 指定了它的长度，它的容量和长度相同。
你可以传入第三个参数来指定不同的容量值，但必须不能比长度值小。
比如 make([]int, 0, 10)
Map: 根据 size 大小来初始化分配内存，不过分配后的 map 长度为 0，如果 size 被忽略了，那么会在初始化分配内存时分配一个小尺寸的内存
Channel: 管道缓冲区依据缓冲区容量被初始化。如果容量为 0 或者忽略容量，管道是没有缓冲区的
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

	//------------------make-------------------------

	var s1 []int
	if s1 == nil {
		fmt.Printf("s1 is nil --> %#v \n ", s1) // []int(nil)
	}
	s2 := make([]int, 3)
	if s2 == nil {
		fmt.Printf("s2 is nil --> %#v \n ", s2)
	} else {
		fmt.Printf("s2 is not nill --> %#v \n ", s2) // []int{0, 0, 0}
	}

	//---------------

	var m1 map[int]string
	if m1 == nil {
		fmt.Printf("m1 is nil --> %#v \n ", m1) //map[int]string(nil)
	}
	m2 := make(map[int]string)
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n ", m2)
	} else {
		fmt.Printf("m2 is not nill --> %#v \n ", m2) // map[int]string{}
	}
	var c1 chan string
	if c1 == nil {
		fmt.Printf("c1 is nil --> %#v \n ", c1) //(chan string)(nil)
	}
	c2 := make(chan string)
	if c2 == nil {
		fmt.Printf("c2 is nil --> %#v \n ", c2)
	} else {
		fmt.Printf("c2 is not nill --> %#v \n ", c2) //(chan string)(0xc420016120)
	}
	//------------------

	slice2 := make([]int, 3)
	fmt.Println(slice2)
	fmt.Printf("%#v \n", slice2) //[]int{0, 0, 0}
	modifySlice(slice2)
	fmt.Printf("%#v \n", slice2) //[]int{1, 0, 0}

	slice3 := []int{0, 0, 0}
	fmt.Println(slice3)
	modifySlice(slice3)
	fmt.Println(slice3)
	//----------------
	map2 := make(map[int]string)
	if map2 == nil {
		fmt.Printf("map2 is nil --> %#v \n ", map2)
	} else {
		fmt.Printf("map2 is not nil --> %#v \n ", map2) //map[int]string{}
	}
	modifyMap(map2)
	fmt.Printf("m2 is not nil --> %#v \n ", map2) // map[int]string{0:"string"}
	chan2 := make(chan string)
	if chan2 == nil {
		fmt.Printf("chan2 is nil --> %#v \n ", chan2)
	} else {
		fmt.Printf("chan2 is not nil --> %#v \n ", chan2)
	}
	go modifyChan(chan2)
	fmt.Printf("chan2 is not nil --> %#v ", <-chan2) //"string"  这里货等待chan生成数据
}

func modifySlice(s []int) {
	s[0] = 1
}

func modifyMap(m map[int]string) {
	m[0] = "string"
}
func modifyChan(c chan string) {
	c <- "string"
}

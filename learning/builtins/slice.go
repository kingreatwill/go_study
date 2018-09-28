// GO语言中的可变长数组的切片(slice);
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// 创建和初始化切片;
func slice01() {
	common.Println("slice.go")
	// 直接初始化切片;
	var s11 = []int{1, 2, 3, 4, 5}
	s12 := []int{1, 2, 3, 4, 5}
	// 切片为某数组的引用;
	var arr21 = []int{1, 2, 3, 4, 5}
	s22 := arr21[:]
	// 某数组的起止下标范围创建新的切片;
	var arr31 = []int{1, 2, 3, 4, 5}
	s32 := arr31[1:3]
	// 某切片的某切片的起始下标范围创建新的切片下标范围创建新的切片;
	var s41 = []int{1, 2, 3, 4, 5}
	s42 := s41[1:3]
	// 某数组的某切片的起始下标范围创建新的切片下标范围创建新的切片(缺省最后一个元素);
	var arr51 = []int{1, 2, 3, 4, 5}
	s52 := arr51[3:]
	// 某切片的某切片的起始下标范围创建新的切片下标范围创建新的切片(缺省最后一个元素);
	var s61 = []int{1, 2, 3, 4, 5}
	s62 := s61[3:]
	// 某数组的某切片的起始下标范围创建新的切片下标范围创建新的切片(缺省第一个元素);
	var arr71 = []int{1, 2, 3, 4, 5}
	s72 := arr71[:4]
	// 某切片的某切片的起始下标范围创建新的切片下标范围创建新的切片(缺省第一个元素);
	var s81 = []int{1, 2, 3, 4, 5}
	s82 := s81[:4]
	// 内置函数初始化切片;
	s91 := make([]string, 4, 50)
	s91[0] = "a"
	s91[1] = "b"
	s91[2] = "c"
	s91[3] = "d"
	// runtime error: index out of range;
	// s91[4] = "e"
	// 打印各个切片的长度及容量;
	fmt.Println(len(s11), cap(s12))
	fmt.Println(len(s22), cap(s22))
	fmt.Println(len(s32), cap(s32))
	fmt.Println(len(s42), cap(s42))
	fmt.Println(len(s52), cap(s52))
	fmt.Println(len(s62), cap(s62))
	fmt.Println(len(s72), cap(s72))
	fmt.Println(len(s82), cap(s82))
	fmt.Println(len(s91), cap(s91))
	fmt.Println(s91)
}

// 空的切片(nil);
func slice02() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil.")
	}
}

// 切片的增长(赋予自身的切片);
func slice03() {
	s := make([]int, 5, 10)
	s = s[:6]
	// runtime error: slice bounds out of range;
	// s = s[:11]
	fmt.Println(len(s))
}

// 切片的增长(追加新的元素);
func slice04() {
	var s []int
	s = append(s, 60)
	s = append(s, 61, 62, 63)
	fmt.Println(s)
}

// 切片的扩大容量(拷贝);
func slice05() {
	s := make([]string, 4, 10)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	// 扩大1倍;
	t := make([]string, len(s), cap(s)*2)
	// 拷贝;
	copy(t, s)
	fmt.Println(len(t), cap(t))
}

// 切片的遍历;
func slice06() {
	s := []int{10, 20, 30, 40}
	// range方式;
	for index, value := range s {
		fmt.Printf("Index: %d  Value: %d\r\n", index, value)
	}
	// 传统的for方式;
	for index := 0; index < len(s); index++ {
		fmt.Printf("Index: %d  Value: %d\r\n", index, s[index])
	}
}

// 多维切片的增长(追加新的元素);
func slice07() {
	s := [][]int{{10}, {20, 21, 22}}
	s[0] = append(s[0], 11)
	fmt.Println(s, len(s), len(s[0]), len(s[1]))
}

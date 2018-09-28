// GO语言中的变量(const);
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// 未使用的常量不会引发编译器错误;
func const01() {
	common.Println("const.go")
	// 分组的形式;
	const (
		LENGTH int = 10
		WIDTH  int = 8
	)
	fmt.Println(LENGTH)
}

// 编译器自动推断常量类型;
func const02() {
	// 省略常量的类型(自动推断);
	const (
		LENGTH = 10
		WIDTH  = 20
	)
	// 不能给常量进行重新赋值;
	// LENGTH = 11
	// 不能读取常量的地址;
	// fmt.Println(&LENGTH, LENGTH)
	fmt.Println(LENGTH)
}

// 一行声明多个常量;
func const03() {
	// 省略常量的类型(自动推断);
	const NAME, SIZE = "carrot", 100
	fmt.Println(NAME)
}

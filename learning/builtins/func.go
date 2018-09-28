// GO语言中的函数(func);
// 不支持默认参数;
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// 单返回值;
func funcSingleReturn(a string) string {
	return a
}

// 命名式多返回值(交换);
func funcMultReturn(x, y string) (string, string) {
	return y, x
}

// 支持可变参数(无返回值);
func funcMultParams(str string, a ...int) {
	fmt.Println("%T, %v\n", str, a)
}

// 闭包(计数器);
func funcSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 函数的使用;
func func01() {
	common.Println("func.go")
	// 单返回值;
	fmt.Println(funcSingleReturn("aaaaaaa"))
	// 命名式多返回值(交换);
	fmt.Println(funcMultReturn("aaaaaaa", "bbbbbbbbb"))
	// 支持可变参数(无返回值);
	funcMultParams("aaaaaaa", 1, 2, 3, 4, 5)
	// 支持匿名函数;
	func(s string) {
		fmt.Println(s)
	}("hello go.")
	// 闭包(计数器);
	nextSeq1 := funcSeq()
	fmt.Println(nextSeq1())
	fmt.Println(nextSeq1())
	fmt.Println(nextSeq1())
	nextSeq2 := funcSeq()
	fmt.Println(nextSeq2())
	fmt.Println(nextSeq2())
	fmt.Println(nextSeq2())
}

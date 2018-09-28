// GO语言中的结构体方法(method);
// 方法与函数的区别: 结构体方法的参数前置, 函数(func)的参数后置;
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// 定义结构体;
type Rectangle struct {
	area, length, width int
}

// 传递值类型的方法;
func (r Rectangle) AreaByValue() int {
	r.area = r.length * r.width
	return r.area
}

// 指针类型的方法;
func (r *Rectangle) AreaByPointer() int {
	r.area = r.length * r.width
	return r.area
}

// 值类型与指针类型的方法;
func method01() {
	common.Println("method.go")
	// 初始化结构体;
	r := Rectangle{0, 4, 3}
	fmt.Println("Rectangle1 area is: ", r.AreaByValue(), r.area)
	fmt.Println("Rectangle2 area is: ", r.AreaByPointer(), r.area)
}

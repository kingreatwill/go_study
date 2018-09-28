// GO语言中的分配内存并初始化一个类型为slice、map、chan的对象(make);
// 返回值为引用;
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// make();
func make01() {
	common.Println("make.go")
	s91 := make([]string, 2, 50)
	s91[0] = "abc"
	s91[1] = "efg"
	fmt.Println(s91)
}

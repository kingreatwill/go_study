// GO语言中的结构体(struct);
// GO语言是没有类的概念, 项目实践中可使用结构体代替;
package builtins

import (
	"fmt"
	"reflect"
)
import "github.com/kingreatwill/go_study/learning/common"

// 结构体定义(普通);
type Vertex struct {
	X int
	Y int
}

// 结构体定义(字段标签);
type User struct {
	Name    string "昵称"
	SexType int    "性别"
}

// 结构体初始化;
func struct01() {
	common.Println("struct.go")
	// 方式1;
	//c := Vertex{5, 10}
	// 方式2(推荐规范);
	c := Vertex{X: 5, Y: 10}
	// 方式3;
	// c := new(Vertex)
	fmt.Println(c)
}

// 字段标签;
func struct02() {
	// 初始化;
	u := User{"Mike", 1}
	// 反射;
	v := reflect.ValueOf(u)
	t := v.Type()
	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}
}

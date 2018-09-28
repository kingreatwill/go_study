// GO语言中的分配内存(new);
// 返回值为指针;
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

type Object struct {
	Id   int
	Name string
}

// new();
func new01() {
	common.Println("new.go")
	// 指针;
	s1 := new(Object)
	s1.Id = 1
	s1.Name = "cat"
	// struct类型的值;
	s2 := Object{Id: 2, Name: "tom"}
	// output: &{1 cat} {2 tom};
	fmt.Println(s1, s2)
}

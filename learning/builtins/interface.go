// GO语言中的接口(interface);
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// 结构体(基类);
type Human struct {
	Name  string
	Age   int
	Phone string
}

// 结构体(学生);
type Student struct {
	// 嵌套结构体;
	Human
	School string
	Loan   float32
}

// 结构体(雇员);
type Employee struct {
	// 嵌套结构体;
	Human
	Company string
	Money   float32
}

// 接口实现;
func (h Human) Living() {
	fmt.Println(h.Name, "Living.")
}

// 接口实现;
func (h Human) Sing(lyrics string) {
	fmt.Println(h.Name, lyrics)
}

// 接口实现;
func (h Human) Belief() string {
	return "The Truth"
}

// 定义接口;
type IMen interface {
	Living()
	Sing(lyrics string)
	Belief() string
}

// 接口的使用;
func interface01() {
	common.Println("interface.go")

	// 初始化;
	student := Student{Human{"mike", 25, "222-222-XXX"}, "MIT", 100.00}
	employee := Employee{Human{"paul", 30, "222-222-XXX"}, "Golang Inc.", 2000.00}

	// 声明接口;
	var i IMen

	// 学生;
	i = student
	i.Living()
	i.Sing("We are the champions.")
	r := i.Belief()
	fmt.Println(r)

	// 雇员;
	i = employee
	i.Living()
	i.Sing("既然为了活命而工作, 为什么又要为了工作而卖命?")
	r = i.Belief()
	fmt.Println(r)
}

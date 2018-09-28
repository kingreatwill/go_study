// GO语言中的异常处理(panic与recover);
// 当一个函数在执行过程中出现了异常或遇到panic(), 正常语句就会立即终止, 然后执行defer语句, 再报告异常信息, 最后退出goroutine;
// 如果在defer中使用了recover()函数, 则会捕获错误信息;
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// Dosomething;
func panicrecoverDosomething() {
	fmt.Println("Do Something")
	panicrecoverThrow("出现错误.")
}

// 抛出异常;
func panicrecoverThrow(error string) {
	panic(error)
}

// 捕获异常;
func panicrecoverLog() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

// 异常处理;
func panicrecover01() {
	common.Println("panicrecover.go")
	// 延迟函数;
	defer panicrecoverLog()
	// Dosomething;
	panicrecoverDosomething()
}

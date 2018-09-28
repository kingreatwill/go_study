// 编译器将根据目录层级自动定位完整包名(书写所在目录名即可);
package builtins

// 包fmt有格式化I/O函数(类似于C语言的printf和scanf);
import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// 大写字母开头的函数才会被导出(外部访问, 相当于public);
func helloWorld() {
	common.Println("helloworld.go")
	// 不使用分号进行结尾;
	fmt.Println("Hello, World!")
}

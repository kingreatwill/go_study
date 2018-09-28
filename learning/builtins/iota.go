// GO语言中的自增量(iota);
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// 优雅的自增量(从iota=0开始);
func iota01() {
	common.Println("iota.go")
	// 分组的形式;
	const (
		LOW = 5 * iota
		MEDIUM
		HIGH
	)
	// 0, 5, 10;
	fmt.Println(LOW, MEDIUM, HIGH)
}

// 中断iota;
func iota02() {
	// 分组的形式;
	const (
		LOW = 5 * iota
		MEDIUM
		HIGH = 100
		SUPER
		BAND = iota
	)
	// 0, 5, 100, 100, 4;
	fmt.Println(LOW, MEDIUM, HIGH, SUPER, BAND)
}

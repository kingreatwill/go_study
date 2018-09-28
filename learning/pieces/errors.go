// GO语言中的异常(errors包);
package pieces

import "fmt"
import "errors"
import "github.com/kingreatwill/go_study/learning/common"

// 创建新异常;
func errors01() {
	common.Println("errors.go")
	// 创建新异常;
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err.Error())
	}
}

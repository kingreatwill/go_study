// 通用抛出异常;
package common

import "fmt"

func PanicIf(e error) {
	if e != nil {
		panic(e)
	}
}

func LoggingIf() {
	if e := recover(); e != nil {
		fmt.Println(e)
	}
}

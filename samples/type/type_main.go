package main

import (
	"fmt"
	"reflect"
)

func main() {
	fd := Field{"nm字段", "ty字段", "rmk", "bind"}
	fdv := reflect.ValueOf(fd)
	fdt := fdv.Type()

	for i, n := 0, fdt.NumField(); i < n; i++ {

		fmt.Printf("%s %s: %v\n", fdt.Field(i).Name, fdt.Field(i).Tag.Get("remark"), fdv.Field(i))
	}
}

//Field 字段;
type Field struct {
	Named   string `remark:"字段"`
	Typed   string `remark:"数值类型"`
	Remark  string `remark:"描述"`
	Binding string `remark:"分类类型(枚举)" binding:"gte=0"`
}

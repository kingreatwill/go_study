/*
GO语言中的变量(variable)
bool
string
int
int8
int16
int32
int64
uint
uint8
uint16
uint32
uint64
uintptr
byte
float32
float64
complex64
complex128
*/
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// 自动类型推断和零初始化;
func var01() {
	common.Println("variable.go")
	// 自动推断a为int;
	var a = 1
	// b初始化为0;
	var b int
	// 支持同时打印多个输出;
	fmt.Println(a, b)
}

// 省略var的简短模式;
func var02() {
	// 声明变量并赋值(相当于C#的int a = 1语法);
	a := 1
	b, s := 2, "go"
	fmt.Println(a, b, s)
}

// 未使用的局部变量视为错误;
func var03() {
	var a, s = 1, "abc"
	// 注意: 是局部变量而不是全局变量;
	// fmt.Println(a)
	fmt.Println(a, s)
}

// 变量声明总结;
func var04() {
	// 类型放在变量的后面(零初始化);
	var a, b, c int
	var d, e byte
	var f, g float32
	// 自动推断a为int;
	var h = 42
	// 混合变量声明;
	var i, j = 100, "abc"
	// 按组的形式;
	var (
		k, l int
		m, n = 100, "abc"
	)
	// 声明变量并赋值(相当于C#的int z = 999语法);
	z := 999
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, z)
}

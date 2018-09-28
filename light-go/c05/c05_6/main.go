package main

import (
	
	"fmt"
)

//测试
//Go 语言中自带有一个轻量级的测试框架 testing 和自带的 go test 命令来实现单元测试和性能测试
//testing 框架和其他语言中的测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基
//于该框架写相应的压力测试用例。通过单元测试，可以解决如下问题:
//1.确保每个函数是可运行，并且运行结果是正确的
//2.确保写出来的代码性能是好的，
//3.单元测试能及时的发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，
//而性能测试的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保持稳定




func main() {
	fmt.Println("hello world")
}

//细节
//1.测试用例文件名必须以 _test.go 结尾。 比如 cal_test.go , cal 不是固定的。
//2.测试用例函数必须以 Test 开头，一般来说就是 Test+被测试的函数名，比如 TestAddUpper
//3.TestAddUpper(t *tesing.T) 的形参类型必须是 *testing.T
//4. 一个测试用例文件中，可以有多个测试用例函数，比如 TestAddUpper、TestSub
//5.运行测试用例指令
// go test [如果运行正确，无日志，错误时，会输出日志]
// go test -v [运行正确或是错误，都输出日志]
//6.当出现错误时，可以使用 t.Fatalf 来格式化输出错误信息，并退出程序
//7.t.Logf 方法可以输出相应的日志
//8.测试用例函数，并没有放在 main 函数中，也执行了，这就是测试用例的方便之处.
//9.PASS 表示测试用例运行成功，FAIL 表示测试用例运行失败
//10.测试单个文件，一定要带上被测试的原文件
//11.go test -v cal_test.go cal.go
//12.测试单个方法 go test -v -test.run TestAddUpper
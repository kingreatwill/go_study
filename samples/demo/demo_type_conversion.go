package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	var i int = 15

	fmt.Printf("%d Uint64 %d\n", i, uint64(i))

	var a string
	a = "123456"
	b, error := strconv.Atoi(a)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
	}
	b = b + 1
	fmt.Println(b)
	var c int = 1234
	d := strconv.Itoa(c) //数字变成字符串
	d = d + "sdfs"
	fmt.Println(d)
	var e interface{}
	e = 10
	switch v := e.(type) {
	case int:
		fmt.Println("整型", v)
		break
	case string:
		fmt.Println("字符串", v)
		break

	}
}

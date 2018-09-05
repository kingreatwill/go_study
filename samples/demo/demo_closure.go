package main

import "fmt"

func main() {
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			i++
			fmt.Printf("i,j:%d,%d\n", i, j)
		}
	}() //将一个无需参数返回值为匿名函数的函数赋值给a()

	a()
	j *= 2
	// i*=2这样是错的
	a()

	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	//---------------带参数的闭包函数调用
	addFunc := add(1, 2)
	fmt.Println(addFunc())
	fmt.Println(addFunc())
	fmt.Println(addFunc())

	add_func := add2(1, 2)
	fmt.Println(add_func(1, 1))
	fmt.Println(add_func(0, 0))
	fmt.Println(add_func(2, 2))

	f := fibonacci1()         //  返回一个闭包函数
	for i := 0; i < 10; i++ { // 检测下前10个值
		fmt.Println(f())
	}
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 闭包使用方法
func add(x1, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x1 + x2
	}
}

// 闭包使用方法
func add2(x1, x2 int) func(x3 int, x4 int) (int, int, int) {
	i := 0
	return func(x3 int, x4 int) (int, int, int) {
		i++
		return i, x1 + x2, x3 + x4
	}
}

func fibonacci1() func() int {
	back1, back2 := -1, 1 // 预先定义好前两个值

	return func() int {
		// 重新赋值(这个就是核心代码)
		back1, back2 = back2, (back1 + back2)
		return back2
	}
}

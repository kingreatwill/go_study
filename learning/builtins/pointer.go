// GO语言中的指针(pointer);
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

func update01(a int) {
	a = a + 100
}

func update02(a *int) {
	*a = *a + 100
}

func update03(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

type Goods struct {
	high   float64
	low    float64
	weight float64
}

func update04(goods *Goods) {
	goods.high = 11.01
	goods.low = 11.02
	goods.weight = 11.03
}

// 按值传递;
func pointer01() {
	common.Println("pointer.go")
	a := 1000
	update01(a)
	fmt.Println("val:", a)
}

// 指针传递;
func pointer02() {
	a := 1000
	update02(&a)
	fmt.Println("val:", a)
}

// 两值交换;
func pointer03() {
	a, b := 1000, 2000
	update03(&a, &b)
	fmt.Println("val:", a, b)
}

// 通过new创建指针;
func pointer04() {
	a := new(int)
	update02(a)
	fmt.Println("val:", *a)
}

// 结构体指针;
func pointer05() {
	goods := Goods{high: 10.01, low: 10.02, weight: 10.03}
	fmt.Println("Original Goods Data:", goods)
	update04(&goods)
	fmt.Println("Modified Goods Data:", goods)
}

// 指针的指针;
func pointer06() {
	a := 6000
	p1 := &a
	p2 := &p1
	fmt.Println(a, *p1, **p2)
}

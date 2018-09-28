// GO语言中的不可变长的数组(array);
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// 一维数组声明(单项下标赋值);
func array01() {
	common.Println("array.go")
	// 单项下标赋值;
	var a [3]int
	a[0] = 1
	a[1] = 2
	// 不赋值则为0;
	// a[2] = 3
	fmt.Println(a[0], a[1], a[2])
}

// 一维数组声明(声明并赋值);
func array02() {
	// 声明并赋值;
	a := [3]int{1, 2}
	// 不赋值则为0;
	// a := [3]int{1, 2, 3}
	fmt.Println(a[0], a[1], a[2])
}

// 一维数组声明(自动推断数组长度);
func array03() {
	// 自动推断数组长度;
	a := [...]int{1, 2, 3}
	fmt.Println(a[0], a[1], a[2])
}

// 一维数组声明(声明并按下标初始化);
func array04() {
	// 声明并按下标初始化;
	a := [3]int{0: 1, 2: 3}
	fmt.Println(a[0], a[1], a[2])
}

// 一维数组声明(数组赋值给数组);
func array05() {
	// 一个数组可以被赋值给任意相同类型的数组;
	var a [5]string
	// 长度不同即使类型相同的数组赋值将编译失败;
	// var b [4]string
	z := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	a = z
	// cannot use z (type [5]string) as type [4]string in assignment;
	// b = z
	fmt.Println(a[0], a[1], a[2], a[3], a[4])
}

// 一维数组声明(指针数组赋值给指针数组);
func array06() {
	var a [3]*string
	// 指针数组的单项下标赋值;
	z := [3]*string{new(string), new(string), new(string)}
	*z[0] = "Red"
	*z[1] = "Blue"
	*z[2] = "Green"
	// 指针数组赋值给指针数组;
	a = z
	// 打印值;
	fmt.Println(*a[0], *a[1], *a[2])
	// 打印地址;
	fmt.Println(a[0], a[1], a[2])
}

// 一维数组声明(遍历);
func array07() {
	// 声明并按下标初始化;
	a := [3]int{0: 1, 2: 3}
	// 遍历;
	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}
}

// 多维数组声明(遍历);
func array08() {
	// 声明并赋值;
	a := [3][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}}
	// 遍历;
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func main() {
	// 我们初始化一个学生
	mark1 := Student{Human{"Mark", 25, 120}, "Computer Science"}
	//mark2 := Student{"Mark", 25, 120, "Computer Science"}
	// 我们访问相应的字段
	fmt.Println("His mark1.name is ", mark1.name)
	fmt.Println("His mark1.Human.name is ", mark1.Human.name)
	mark1.Human.name = "M2"
	fmt.Println("His mark1.name is ", mark1.name)
	fmt.Println("His mark1.Human.name is ", mark1.Human.name)
	mark1.name = "M3"
	fmt.Println("His mark1.name is ", mark1.name)
	fmt.Println("His mark1.Human.name is ", mark1.Human.name)
}

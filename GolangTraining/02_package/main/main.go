package main

import (
	"fmt"

	"github.com/kingreatwill/go_study/GolangTraining/02_package/icomefromalaska"
	"github.com/kingreatwill/go_study/GolangTraining/02_package/stringutil"
	//someAlias "github.com/kingreatwill/go_study/GolangTraining/02_package/icomefromalaska"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(winniepooh.BearName)

	s := "hello你好"
	fmt.Println(len(s)) //输出长度为11

	r := []rune(s)
	fmt.Println(r)      //[104 101 108 108 111 20320 22909]
	fmt.Println(len(r)) //输出长度为7
	s = "你好"
	fmt.Println(len(s))         //输出长度为6
	fmt.Println(len([]rune(s))) //输出长度为2
	s = "你"
	fmt.Println([]byte(s))      //输出[228 189 160]
	fmt.Println(len([]byte(s))) //输出3
	fmt.Println(len([]rune(s))) //输出1
	fmt.Println(rune('你'))      //输出20320

}

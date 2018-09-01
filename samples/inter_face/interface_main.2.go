package main

import (
	"fmt"
	"reflect"
)

type I interface {
	M()
}
type T1 struct{}

func (T1) M() { fmt.Println("T1.M") }

type T2 struct{}

func (T2) M() { fmt.Println("T2.M") }

func f(i I) {
	i.M()
	fmt.Println(reflect.TypeOf(i).PkgPath(), reflect.TypeOf(i).Name())
	fmt.Println(reflect.TypeOf(i).String())
}

func main() {
	f(T1{}) // "T1.M"
	f(T2{}) // "T2.M"

}
